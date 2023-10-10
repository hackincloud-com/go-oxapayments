package oxapayments

import (
	"encoding/json"
	"io"

	httpclient "github.com/Hax0rCompany/repweb/pkg/repclient"
	"github.com/spf13/cobra"
)

// Start an Oxapayments with using configuration that already init
// use invoice,createpayment,paymentinfo
func (oxa *OxaConfig) Start(reqConfig *OxaReqs) OxaResp {
	body := reqConfig.Requests()
	o := OxaResp{}
	o.Invoice(body)
	o.CreatePayment(body)
	o.PaymentInfo(body)
	return o
}

// Requests with reqconfig struct
// return body
func (reqConfig *OxaReqs) Requests() []byte {
	client := &httpclient.ReqStruct{
		Url:     reqConfig.Url,
		Method:  "POST",
		Body:    reqConfig.Body,
		Timeout: 60,
	}

	res := client.DoRequests()
	if res.Error != nil {
		cobra.CheckErr(res.Error)
	}

	body, _ := io.ReadAll(res.Response.Body)
	return body
}

// create invoice
// please initiate OxaResp first before use it
func (o *OxaResp) Invoice(b []byte) {
	Invoice := &RespInvoice{}
	if err := json.Unmarshal(b, Invoice); err != nil {
		cobra.CheckErr(err)
	}
	o.RespInvoice = *Invoice
}

// create payment
// please initiate OxaResp first before use it
func (o *OxaResp) CreatePayment(b []byte) {
	Payment := &RespCreatePayment{}
	if err := json.Unmarshal(b, Payment); err != nil {
		cobra.CheckErr(err)
	}
	o.RespCreatePayment = *Payment
}

// get payment info with trackId config
// please initiate OxaResp first before use it
func (o *OxaResp) PaymentInfo(b []byte) {
	Info := &RespPaymentInfo{}
	if err := json.Unmarshal(b, Info); err != nil {
		cobra.CheckErr(err)
	}
	o.RespPaymentInfo = *Info
}
