package main

import (
	"Crunch-Garage/go-stripe/config"
	"Crunch-Garage/go-stripe/models"
	"fmt"

	//"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/customer"
)

func main() {
	//cus_MH2pXcoeax4Tcl

	/*GLOBALY SET API KEY*/
	/*initialize stripe client globally if using one stripe account*/
	stripe.Key = config.EnvStripeApiKey()
	//c, _ := customer.Get("cus_4eC39HqLyjvKYx", nil)

	/*SET API KEY PER REQUEST*/
	/* 1- if using multiple stripe accounts say one in USA and another in Europe
	   2- set an api key each time when making a new api per  request
	   3- if building a platform for multiple vendors i.e UBER, LYFT   pass the api key in requests
	*/
	//sc := &client.API{}
	//sc.Init(config.EnvStripeApiKey(), nil)
	//c, _ := sc.Customers.Get("cus_MH2lQHxvjab6TA", nil)

	/*get charge params from api request*/
	chargeParams := &models.StripeCharge{
		Amount:      505,
		Description: "Food delivery",
	}

	/*get customer params from api request*/
	customerParams := &models.StripeCustomer{
		Name:        "John doe",
		Phone:       "+254719686126",
		Email:       "johndoe2@example.com",
		Description: "Test Customer",
	}

	/*initiate creation of charge*/
	c, _ := CreateStripeCharge(customerParams, chargeParams)
	fmt.Println(c)

	/*Get Stripe Customer*/
	//c := GetStripeCutomer("cus_MH33H53kzlxe4y")
	//fmt.Println(c)

}

/*CREATE NEW CUSTOMER*/
func CreateStripeCustomer(customerParams *models.StripeCustomer) (*stripe.Customer, error) {

	params := &stripe.CustomerParams{
		Name:        stripe.String(customerParams.Name),
		Phone:       stripe.String(customerParams.Phone),
		Email:       stripe.String(customerParams.Email),
		Description: stripe.String(customerParams.Description),
	}

	newCustomer, err := customer.New(params)

	if err != nil {
		return nil, err
	}

	return newCustomer, err

}

/*GET STRIPE CUSTOMER*/
func GetStripeCutomer(customer_id string) (*stripe.Customer, error) {

	//params:= &stripe.CustomerParams{}
	//params.SetStripeAccount("acc_")  // pass related account on acc_
	c, err := customer.Get(customer_id, nil) // if passing param repalce nil with params

	if err != nil {
		return nil, err
	}

	return c, nil
}

/*CREATE STRIPE CHARGE*/
func CreateStripeCharge(customerParams *models.StripeCustomer, chargeParams *models.StripeCharge) (*stripe.Charge, error) {

	newCharge := &stripe.ChargeParams{
		Amount:   stripe.Int64(chargeParams.Amount),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		//Customer:     &newCustomer.ID,
		ReceiptEmail: stripe.String(customerParams.Email),
		Description:  stripe.String(chargeParams.Description),
		Source:       &stripe.SourceParams{Token: stripe.String("tok_visa")},
	}

	customerCharge, chargeError := charge.New(newCharge)

	if chargeError != nil {
		return nil, chargeError
	}

	return customerCharge, nil
}
