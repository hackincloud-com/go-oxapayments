# Simple OXAPayments Go Library 
## ( NON-KYC AND LOW-FEE Crypto Gateaway )

Register first in oxapayments and create ur apikey in merchant

## Installation

```bash
go get github.com/hackincloud-com/go-oxapayments
```

## Usage

Just load the config with all the credentials and config from `InitConfig` then add your configurations using these functions

### Basic Usage for Merchant Create Invoice
```go
package main

import (
	"fmt"
	"log"
	"strings"
	oxapayments"github.com/hackincloud-com/go-oxapayments"
)

var payment = oxapayments.InitConfig(
		oxapayments.WithCurrency("USDT"),
		oxapayments.WithApiKey("sandbox"),
)


func main() {
	// Set an value for payment
	payment.OrderId = "ORDER-1"
	payment.Description = "Order For Mango"
	payment.Email = "i_like_mango@client.com"
	payment.Amount = 1 
    // Set Merchants api to start invoice
	req := payment.Set("https://api.oxapay.com/merchants/request")
	results := payment.Start(req)
	// see the examples func below
	CheckPaymentInfo(results.RespInvoice.TrackID) 
    fmt.Println(results.PayLink) // get the paylink
}
```

### Check Payment Info
```go
func CheckPaymentInfo(IdFromInvoice string){
	// Track Transaction status and more
    payment.TrackID = IdFromInvoice
    getinfo := payment.Set("https://api.oxapay.com/merchants/inquiry")
	res := payment.Start(getinfo)
    fmt.Println(res.Status) // status trx
}
```

## CallBack ( Tract Transaction )
The URL where payment information will be sent. Use this to receive notifications about the payment status.
( disclaimer : im using fiber )
#### Please check "callback.go" files.