package stripefunctions

import (
	"Crunch-Garage/go-stripe/models"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
)

/*CREATE NEW CUSTOMER*/
func CreateStripeCustomer(customerParams *models.StripeCustomer) (*stripe.Customer, error) {

	params := &stripe.CustomerParams{
		Name:        stripe.String(customerParams.Name),
		Phone:       stripe.String(customerParams.Phone),
		Email:       stripe.String(customerParams.Email),
		Description: stripe.String(customerParams.Description),
	}

	newCustomer, err := customer.New(params)
	/*create payment method*/
	// createPaymethod,createPaymethodErr  := CreatePaymentMethod("cus_MH4aPLRtAMUd6o", cardParams)

	if err != nil {
		return nil, err
	}

	return newCustomer, err

}

/*GET STRIPE CUSTOMER*/
func GetStripeCustomer(customer_id string) (*stripe.Customer, error) {

	//params:= &stripe.CustomerParams{}
	//params.SetStripeAccount("acc_")  // pass related account on acc_
	c, err := customer.Get(customer_id, nil) // if passing param repalce nil with params

	if err != nil {
		return nil, err
	}

	return c, nil
}
