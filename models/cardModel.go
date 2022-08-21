package models

type StripeCard struct {
	Number   string `json:"number"`
	ExpMonth string `json:"expMonth"`
	ExpYear  string `json:"expYear"`
	CVC      string `json:"cVC"`
	Type     string `json:"type"`
}
