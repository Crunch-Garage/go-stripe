package models

type StripeCharge struct {
	Amount       int64  `json:"amount"`
	ReceiptEmail string `json:"receiptEmail"`
	Description  string `json:"description"`
}
