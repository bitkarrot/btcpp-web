package getters

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/base58btc/btcpp-web/internal/config"
	"github.com/base58btc/btcpp-web/internal/types"
	"github.com/niftynei/cln-checkout/checkout"
)

type InvoiceDesc struct {
	Desc    string
	Email   string
	ConfRef string
	TixType string
	AmtUSD  uint64
}

var descMatch *regexp.Regexp = regexp.MustCompile("(?P<Desc>.*) \\(\\$(?P<AmtUSD>[0-9]+)USD\\) \\[(?P<ConfRef>[a-z0-9-]+)\\|(?P<Email>.*)\\|(?P<TixType>[a-z]+)\\]")

func MakeCLNDesc(conf *types.Conf, email string, tixPrice uint, isLocal bool) string {
	tixType := "genpop"
	if isLocal {
		tixType = "local"
	}
	return fmt.Sprintf("1 ticket for the %s ($%dUSD) [%s|%s|%s]", conf.Desc, tixPrice, conf.Ref, email, tixType)
}

func parseCLNDesc(desc string) (*InvoiceDesc, error) {

	matches := descMatch.FindStringSubmatch(desc)

	if len(matches) != 6 {
		return nil, fmt.Errorf("Desc didn't match expected pattern %s", desc)
	}

	amtUSD, err := strconv.ParseUint(matches[descMatch.SubexpIndex("AmtUSD")], 10, 64)
	if err != nil {
		return nil, err
	}

	return &InvoiceDesc{
		Desc:    matches[descMatch.SubexpIndex("Desc")],
		ConfRef: matches[descMatch.SubexpIndex("ConfRef")],
		Email:   matches[descMatch.SubexpIndex("Email")],
		TixType: matches[descMatch.SubexpIndex("TixType")],
		AmtUSD:  amtUSD,
	}, nil
}

func HandleCLNInvoiceEvent(ctx *config.AppContext, invoice *checkout.InvoiceEvent) bool {
	invDesc, err := parseCLNDesc(invoice.Description)
	if err != nil {
		ctx.Err.Printf("Unable to parse invoice desc %s", err)
		return true
	}

	ctx.Infos.Printf("%s invoice. (label: %s)", invoice.Status, invoice.Label)
	if invoice.Status != "paid" {
		return true
	}

	if invoice.PaidAt == nil {
		ctx.Err.Printf("Paid invoice missing paidat?? %v", invoice)
		return true
	}

	entry := types.Entry{
		ID:       invoice.Label,
		ConfRef:  invDesc.ConfRef,
		Total:    int64(invDesc.AmtUSD * 100),
		Currency: "USD",
		Created:  *invoice.PaidAt,
		Email:    invDesc.Email,
		Items: []types.Item{
			{
				Total: int64(invDesc.AmtUSD * 100),
				Desc:  invDesc.Desc,
				Type:  invDesc.TixType,
			},
		},
	}

	added, err := AddTickets(ctx, ctx.Notion, &entry, "CLN")
	if err != nil {
		ctx.Err.Printf("!!! Unable to add tickets %s: %v", err, entry)
		return false
	}

	if added > 0 {
		ctx.Infos.Printf("Added %d tickets! (%s)", added, entry.ID)
	}
	return true
}
