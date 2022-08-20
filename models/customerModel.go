package models

type StripeCustomer struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Description string `json:"description"`
}
