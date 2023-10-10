package oxapayments

import "io"

type OxaConfig struct {
	MerchantApiKey string `json:"merchant"`
	TrackID        string `json:"trackId"`
	Currency       string `json:"currency"`
	PayCurrency    string `json:"payCurrency"`
	Lifetime       int    `json:"lifeTime"`
	FeePaidByPayer int    `json:"feePaidByPayer"`
	UnderPaidCover int64  `json:"underPaidCover"`
	CallbackUrl    string `json:"callbackUrl"`
	ReturnUrl      string `json:"returnUrl"`
	Description    string `json:"description"`
	OrderId        string `json:"orderId"`
	Amount         int64  `json:"amount"`
	Email          string `json:"email"`
}

type OxaReqs struct {
	Url  string
	Body io.Reader
}

type OxaResp struct {
	RespInvoice       RespInvoice
	RespCreatePayment RespCreatePayment
	RespPaymentInfo   RespPaymentInfo
}

// RespInvoice is struct for Respond API Invoice
type RespInvoice struct {
	Result    int    `json:"result"`
	Message   string `json:"message"`
	TrackID   string `json:"trackId"`
	ExpiredAt int64  `json:"expiredAt"`
	PayLink   string `json:"payLink"`
}
type RespCreatePayment struct {
	Result         int    `json:"result"`
	Message        string `json:"message"`
	TrackID        string `json:"trackId"`
	Amount         string `json:"amount"`
	Currency       string `json:"currency"`
	PayAmount      string `json:"payAmount"`
	PayCurrency    string `json:"payCurrency"`
	Network        string `json:"network"`
	Address        string `json:"address"`
	Rate           string `json:"rate"`
	ExpiredAt      int64  `json:"expiredAt"`
	CreatedAt      string `json:"createdAt"`
	LifeTime       int    `json:"lifeTime"`
	UnderPaidCover int    `json:"underPaidCover"`
	FeePaidByPayer int    `json:"feePaidByPayer"`
	OrderID        string `json:"orderId"`
	Email          string `json:"email"`
	Description    string `json:"description"`
	CallbackURL    string `json:"callbackUrl"`
	QRCode         string `json:"QRCode"`
}
type RespPaymentInfo struct {
	Result         int    `json:"result"`
	Message        string `json:"message"`
	TrackID        string `json:"trackId"`
	Address        string `json:"address"`
	Amount         string `json:"amount"`
	Currency       string `json:"currency"`
	PayAmount      string `json:"payAmount"`
	PayCurrency    string `json:"payCurrency"`
	Network        string `json:"network"`
	FeePaidByPayer int    `json:"feePaidByPayer"`
	UnderPaidCover int    `json:"underPaidCover"`
	Status         string `json:"status"`
	Type           string `json:"type"`
	TxID           string `json:"txID"`
	Date           string `json:"date"`
	PayDate        string `json:"payDate"`
	Email          string `json:"email"`
	OrderID        string `json:"orderId"`
	Description    string `json:"description"`
}

type RespPaymentHistory struct {
}
type OxaConfFunc func(*OxaConfig)
