package stripefunctions

import (
	"Crunch-Garage/go-stripe/models"
	"fmt"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentmethod"
)

/*CREATE PAYMENT METHOD*/
func CreatePaymentMethod(customer_id string, card *models.StripeCard) (*stripe.PaymentMethod, error) {

	params := &stripe.PaymentMethodParams{
		Card: &stripe.PaymentMethodCardParams{
			Number:   stripe.String(card.Number),
			ExpMonth: stripe.String(card.ExpMonth),
			ExpYear:  stripe.String(card.ExpYear),
			CVC:      stripe.String(card.CVC),
		},
		Type: stripe.String("card"),
	}

	paymentmethod, err := paymentmethod.New(params)

	if err != nil {
		return nil, err
	}

	_, attachment_error := AttachCustomerPaymentMethod(customer_id, paymentmethod.ID)

	/*if there is an error with attaching the payment method tou customer return error*/
	if attachment_error != nil {
		return nil, attachment_error
	}

	return paymentmethod, nil
}

/*LINK PAYMENT METHOD TO CUSTOMER*/
func AttachCustomerPaymentMethod(customer_id string, paymentMethod_id string) (*stripe.PaymentMethod, error) {

	params := &stripe.PaymentMethodAttachParams{
		Customer: stripe.String(customer_id),
	}
	paymentmethod, err := paymentmethod.Attach(
		paymentMethod_id,
		params,
	)

	if err != nil {
		return nil, err
	}

	fmt.Println(paymentmethod)

	return paymentmethod, nil
}
