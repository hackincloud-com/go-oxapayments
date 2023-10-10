package main

import (
	"fmt"
	"log"

	"github.com/hackincloud-com/go-oxapayments"
)

// If you using Whitelist Payment you need setup the oxapayments.WithPayCurrency()
var payment = oxapayments.InitConfig(
	oxapayments.WithCurrency("USDT"),
	oxapayments.WithApiKey("sandbox"),
)

// Examples of specified without .Start
func Specified() {
	results := &oxapayments.OxaResp{}
	set := payment.Set("https://api.oxapay.com/merchants/request")
	req := set.Requests()
	results.Invoice(req)                   // get Invoice
	fmt.Printf("%+v", results.RespInvoice) // here are the results
}

func main() {
	created := CreatePayment() // check below for this function
	// because in this examples i use merchant invoice i set api to this
	set := created.Set("https://api.oxapay.com/merchants/request")
	// you can directly start or be speficied like in function Specified()
	results := payment.Start(set) // remember using payment ( oxapayments Config )
	// you can access it oxapayments.OxaResp

	if results.RespInvoice.Result != 100 {
		log.Fatalf("the results is not success %+v\n", results.RespInvoice)
	}
	fmt.Printf("Results : %+v\n", results.RespInvoice.PayLink)
	CheckPaymentInfo(results.RespInvoice.TrackID)
}

// Create Payments Configuration
func CreatePayment() oxapayments.OxaConfig {
	// set payment
	payment.OrderId = "ORDER-1-MANGO"
	payment.Description = "Ordering Mango"
	payment.Email = "i_like_mango@client.com"
	payment.Amount = 10 // it will charge 10 USDT
	return payment
}

func CheckPaymentInfo(trackId string) {
	payment.TrackID = trackId
	getinfo := payment.Set("https://api.oxapay.com/merchants/inquiry")
	results := payment.Start(getinfo)
	fmt.Printf("Payment Info : %+v\nTrackID : %+v \n", results.RespPaymentInfo.Status, results.RespPaymentInfo.TrackID)
}
