package getters

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"time"
)

const COINDESK_ENDPOINT = "https://api.coindesk.com/v1/bpi/currentprice/USD.json"

/*
 Example response from endpoint
{
  "time": {
    "updated": "Nov 30, 2023 23:12:00 UTC",
    "updatedISO": "2023-11-30T23:12:00+00:00",
    "updateduk": "Nov 30, 2023 at 23:12 GMT"
  },
  "disclaimer": "This data was produced from the CoinDesk Bitcoin Price Index (USD). Non-USD currency data converted using hourly conversion rate from openexchangerates.org",
  "bpi": {
    "USD": {
      "code": "USD",
      "rate": "37,630.0755",
      "description": "United States Dollar",
      "rate_float": 37630.0755
    }
  }
}
*/

var lastUpdate *CoindeskPrice

type CoindeskPrice struct {
	UpdatedAt time.Time
	USDCents  uint64
}

func (cp *CoindeskPrice) UnmarshalJSON(data []byte) error {
	type TimeObj struct {
		Updated    string `json:"updated"`
		UpdatedISO string `json:"updatedISO"`
		UpdatedUK  string `json:"updateduk"`
	}
	type Currency struct {
		Code        string  `json:"code"`
		Rate        string  `json:"rate"`
		Description string  `json:"description"`
		RateFloat   float64 `json:"rate_float"`
	}
	type Bpi struct {
		USD Currency `json:"USD"`
	}
	var coindeskResp struct {
		Time       TimeObj `json:"time"`
		Disclaimer string  `json:"disclaimer"`
		Bpi        Bpi     `json:"bpi"`
	}

	if err := json.Unmarshal(data, &coindeskResp); err != nil {
		return err
	}

	layout := "2006-01-02T15:04:05+00:00"
	t, err := time.Parse(layout, coindeskResp.Time.UpdatedISO)
	if err != nil {
		return err
	}

	cents := math.Ceil(coindeskResp.Bpi.USD.RateFloat * 100)

	*cp = CoindeskPrice{
		UpdatedAt: t,
		USDCents:  uint64(cents),
	}
	return nil
}

func ConvertToSats(amountCents uint64) (uint64, error) {
	var err error
	if lastUpdate == nil {
		lastUpdate, err = FetchCoindeskPrice()
		if err != nil {
			return 0, err
		}
	}

	now := time.Now().UTC()
	if lastUpdate.UpdatedAt.Before(now.Add(-30 * time.Minute)) {
		update, err := FetchCoindeskPrice()
		if err != nil {
			return 0, err
		}
		lastUpdate = update
	}

	/* How much bitcoin is this, in sats? */
	sats := math.Ceil(float64(amountCents*100000000) / float64(lastUpdate.USDCents))
	return uint64(sats), err
}

func FetchCoindeskPrice() (*CoindeskPrice, error) {
	req, _ := http.NewRequest("GET", COINDESK_ENDPOINT, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error returned from coindesk %d: %s", resp.StatusCode, body)
	}
	var priceUpdate CoindeskPrice
	json.Unmarshal(body, &priceUpdate)
	return &priceUpdate, nil
}
