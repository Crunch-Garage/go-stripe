package models

type StripeCard struct {
	ReceiptEmail string `json:"receiptEmail"`
	Number       string `json:"number"`
	ExpMonth     string `json:"expMonth"`
	ExpYear      string `json:"expYear"`
	CVC          string `json:"cVC"`
}
