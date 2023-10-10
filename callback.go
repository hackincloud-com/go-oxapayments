// Examples of using callback in the Fiber Golang

package oxapayments

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/Hax0rCompany/repweb/internal/app/types"
	"github.com/Hax0rCompany/repweb/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type CallBack struct {
	Status         string `json:"status"`
	TrackID        string `json:"trackId"`
	Amount         string `json:"amount"`
	Currency       string `json:"currency"`
	FeePaidByPayer int    `json:"feePaidByPayer"`
	UnderPaidCover int    `json:"underPaidCover"`
	Email          string `json:"email"`
	OrderID        string `json:"orderId"`
	Description    string `json:"description"`
	Date           string `json:"date"`
	PayDate        int    `json:"payDate"`
	Type           string `json:"type"`
	PayAmount      string `json:"payAmount"`
	PayCurrency    string `json:"payCurrency"`
	Price          string `json:"price"`
}

func Callback(c *fiber.Ctx) error {
	m := new(CallBack)
	if err := utils.ParseBodyAndValidate(c, m); err != nil {
		return c.JSON(&types.MsgResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}
	if m.Type != "payment" {
		return c.JSON(&types.MsgResponse{
			Status:  "error",
			Message: "Invalid Data type",
		})
	}
	hmacHeader := c.Get("HMAC")
	api_secret_key := []byte(viper.GetString("settings.OXAPAY_MERCHANT_API_KEY")) // set vipet.getstring "urapikey" or if u want demo or sandbox use "sandbox"
	hash := hmac.New(sha512.New, api_secret_key)
	hash.Write(c.Body())
	calculatedHMAC := hash.Sum(nil)
	calculatedHMACString := hex.EncodeToString(calculatedHMAC)
	if calculatedHMACString == hmacHeader {
		fmt.Printf("%s\n", string(c.Body()))
		return c.JSON(m)
	} else {
		log.Fatalf("\n invalid hmac\n ur hmac: %s\n encoded hmac: %s \nand byte hmac to str manual %s", hmacHeader, calculatedHMACString, string(calculatedHMAC))
		return fiber.ErrBadGateway
	}
}

// decleare a routes into the callback functions to use it
