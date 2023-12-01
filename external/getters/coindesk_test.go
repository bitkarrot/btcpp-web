package getters

import (
	"math"
	"testing"
)

func TestFetch(t *testing.T) {
	price, err := FetchCoindeskPrice()

	if err != nil {
		t.Fatalf("oops %s", err)
	}

	if price.USDCents == 0 {
		t.Fatalf("oh no %v", price)
	}
}

func TestConvert(t *testing.T) {
	price, err := FetchCoindeskPrice()
	if err != nil {
		t.Fatalf("oops %s", err)
	}

	/* convert 100k to BtC */
	amt := uint64(100000 * 100)
	sats, err := ConvertToSats(amt)
	if err != nil {
		t.Fatalf("oops %s", err)
	}

	usdcents := math.Floor(float64(sats*price.USDCents) / 100000000)

	if amt != uint64(usdcents) {
		t.Fatalf("amt: %d, usd: %d", amt, uint(usdcents))
	}
}
