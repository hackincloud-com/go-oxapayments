package oxapayments

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/spf13/cobra"
)

// Initialize Config with the func of the change
// the default config is in the defaultOpts
func InitConfig(opts ...OxaConfFunc) OxaConfig {
	defaultOpts := OxaConfig{
		Currency:       "BTC", // default BTC currency
		FeePaidByPayer: 1,     // default true
		// CallbackUrl:    "https://4d5b-85-203-21-16.ngrok-free.app/auth/callback_oxapay",
		UnderPaidCover: 11, // default 11%
	}
	for _, fn := range opts {
		fn(&defaultOpts)
	}
	if defaultOpts.MerchantApiKey == "" {
		log.Fatalln("you need input a merchant apikey from oxapay.com")
	}
	return defaultOpts
}

// Set an callback url to notify ( if you use localhost please fordward it to NGROK )
func WithCallBackUrl(url string) OxaConfFunc {
	return func(opts *OxaConfig) {
		opts.CallbackUrl = url
	}
}

// Set an client email
func WithEmail(email string) OxaConfFunc {
	return func(opts *OxaConfig) {
		opts.Email = email
	}
}

// Set Your Merchant APIKEY *Important
func WithApiKey(key string) OxaConfFunc {
	return func(opts *OxaConfig) {
		opts.MerchantApiKey = key
	}
}

// Set an api endpoints and Marshalling a Configurations to send in the Oxa Api
// And Check error fatal
// Return an Configuration OxaReqs ( for start )
func (oxa *OxaConfig) Set(url string) *OxaReqs {
	b, err := json.Marshal(oxa)
	cobra.CheckErr(err)
	return &OxaReqs{Url: url, Body: bytes.NewReader(b)}
}

// Set PayCurrency for Create Payment ( NOT INVOICE )
// Importent if you using create payemnt
// Check https://docs.oxapay.com/api-reference/supported-currencies
// E.g WithPayCurrency("TRX")
// and WithCurrency("USD")
func WithPayCurrency(paycurrency string) OxaConfFunc {
	return func(opts *OxaConfig) {
		opts.PayCurrency = paycurrency
	}
}

// Set Currency ( Work in invoice and create payment )
// Currency can be trx,btc
func WithCurrency(currency string) OxaConfFunc {
	return func(opts *OxaConfig) {
		opts.Currency = currency
	}
}

// Set amount int64
// if you set currency btc ind amount 1
// it will be 1 btc
func WithAmount(amount int64) OxaConfFunc {
	return func(opts *OxaConfig) {
		opts.Amount = amount
	}
}

// Set an Order ID
func WithOrderID(id string) OxaConfFunc {
	return func(opts *OxaConfig) {
		opts.OrderId = id
	}
}

// Set an Description
func WithDescription(desc string) OxaConfFunc {
	return func(opts *OxaConfig) {
		opts.Description = desc
	}
}
