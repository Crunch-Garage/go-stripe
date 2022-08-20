package stripefunctions

import (
	"Crunch-Garage/go-stripe/models"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

/*CREATE STRIPE CHARGE*/
func CreateStripeCharge(customerParams *models.StripeCustomer, chargeParams *models.StripeCharge) (*stripe.Charge, error) {

	params := &stripe.ChargeParams{
		Amount:   stripe.Int64(chargeParams.Amount),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		//Customer:     &newCustomer.ID,
		ReceiptEmail: stripe.String(customerParams.Email),
		Description:  stripe.String(chargeParams.Description),
		Source:       &stripe.SourceParams{Token: stripe.String("tok_visa")},
	}

	customerCharge, err := charge.New(params)

	if err != nil {
		return nil, err
	}

	return customerCharge, nil
}
