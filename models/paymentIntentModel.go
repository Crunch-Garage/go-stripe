package models

type StripePaymentIntent struct {
	Amount       int64  `json:"amount"`
	PaymentMethodTypes string `json:"paymentMethodTypes"`
}
