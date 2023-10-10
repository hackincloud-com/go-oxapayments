package oxapayments

import (
	"testing"

	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
)

var (
	Trackid string
)

var payment = InitConfig(
	WithOrderID("TEST-1234"),
	WithAmount(1),
	WithDescription("Testing"),
	WithCurrency("USDT"),
	// for white label payment
	// WithCurrency("USD"),
	// WithPayCurrency("TRX"),
	WithEmail("client@gmail.com"),
	WithApiKey("sandbox"),
)

func TestInvoice(t *testing.T) {
	gologger.DefaultLogger.SetMaxLevel(levels.LevelDebug)
	req := payment.Set("https://api.oxapay.com/merchants/request")
	results := payment.Start(req)
	if results.RespInvoice.Result != 100 {
		gologger.Debug().TimeStamp().Msgf("%+v", results)
		t.Errorf("The status code for fetching api Invoice is %d instead of %d check https://docs.oxapay.com/api-reference/result-code-table for more", results.RespInvoice.Result, 100)
	}

	Trackid = results.RespInvoice.TrackID

	gologger.Debug().Msg(results.RespInvoice.PayLink)
}

func TestPaymentInfo(t *testing.T) {
	gologger.DefaultLogger.SetMaxLevel(levels.LevelDebug)
	payment.TrackID = "89417437"
	results := &OxaResp{}
	getinfo := payment.Set("https://api.oxapay.com/merchants/inquiry")
	req := getinfo.Requests()
	results.PaymentInfo(req)
	gologger.Debug().Msgf("status code : %+v\n", results.RespPaymentInfo)
	if results.RespPaymentInfo.Result != 100 {
		gologger.Debug().TimeStamp().Msgf("%+v", results)
		t.Errorf("The status code for fetching api Invoice is %d instead of %d check https://docs.oxapay.com/api-reference/result-code-table for more", results.RespPaymentInfo.Result, 100)
	}
}
