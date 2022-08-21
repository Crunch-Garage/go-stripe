package stripefunctions

import (
	"Crunch-Garage/go-stripe/models"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

/*CREATE STRIPE CHARGE*/
func CreateStripeCharge(chargeParams *models.StripeCharge) (*stripe.Charge, error) {

	//TODO: for development / testing mode 'tok_visa' is used in place of Token
	//For live enviroment replace 'tok_visa' with 'token' for Token
	params := &stripe.ChargeParams{
		Amount:   stripe.Int64(chargeParams.Amount),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		//Customer:     stripe.String(chargeParams.Customer),
		ReceiptEmail: stripe.String(chargeParams.ReceiptEmail),
		Description:  stripe.String(chargeParams.Description),
		Source:       &stripe.SourceParams{Token: stripe.String(chargeParams.Token)},
	}

	customerCharge, err := charge.New(params)

	if err != nil {
		return nil, err
	}

	return customerCharge, nil
}
