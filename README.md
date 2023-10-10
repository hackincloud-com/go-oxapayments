# Simple OXAPayments Go Library 
## ( NON-KYC AND LOW-FEE Crypto Gateaway )

Register first in oxapayments and create ur apikey in merchant

## Installation

```bash
go get github.com/t101804/oxapayments
```

## Usage

Just load the config with all the credentials and config from `InitConfig` then add your configurations using these functions

```go
package main

import (
	"fmt"
	"log"
	"strings"
	"github.com/t101804/oxapayments"
)

func CheckPaymentInfo(){
     
}
func main() {
    payment := oxapayments.InitConfig(
		oxapayments.WithOrderID("TEST-1234"),
		oxapayments.WithAmount(1), 
		oxapayments.WithDescription("Testing"),
		oxapayments.WithCurrency("USDT"),
		oxapayments.WithEmail("client@gmail.com"),
		oxapayments.WithApiKey("sandbox"),
	)
    // Merchants api is to start invoice
	req := payment.Set("https://api.oxapay.com/merchants/request")
	results := payment.Start(req)
    fmt.Println(results.PayLink) // get the paylink
    // Track Transaction status and more
    payment.TrackID = results.TrackID
    getinfo := payment.Set("https://api.oxapay.com/merchants/inquiry")
	res := payment.Start(getinfo)
    fmt.Println(res.Status) // status trx
}
```

## CallBack ( Tract Transaction )
The URL where payment information will be sent. Use this to receive notifications about the payment status.
( disclaimer : im using fiber )
#### Please check "callback.go" files.