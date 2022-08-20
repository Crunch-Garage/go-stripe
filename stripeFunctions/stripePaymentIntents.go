package stripefunctions

import (
	"Crunch-Garage/go-stripe/models"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

func CreateStripePayIntent(paymentParams *models.StripePaymentIntent) (*stripe.PaymentIntent, error) {

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(paymentParams.Amount),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		PaymentMethodTypes: stripe.StringSlice([]string{
			paymentParams.PaymentMethodTypes,
		}),
	}

	paymentIntent, err := paymentintent.New(params)

	if err != nil {
		return nil, err
	}

	return paymentIntent, nil
}
