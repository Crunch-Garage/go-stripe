package main

import (
	"Crunch-Garage/go-stripe/config"
	"Crunch-Garage/go-stripe/models"
	stripefunctions "Crunch-Garage/go-stripe/stripeFunctions"
	"fmt"

	"github.com/stripe/stripe-go"
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
	// chargeParams := &models.StripeCharge{
	// 	Amount:      505,
	// 	Description: "Food delivery",
	// }

	/*get customer params from api request*/
	// customerParams := &models.StripeCustomer{
	// 	Name:        "John doe",
	// 	Phone:       "+254719686126",
	// 	Email:       "johndoe2@example.com",
	// 	Description: "Test Customer",
	// }

	/*get card details from api request*/
	cardParams := &models.StripeCard{
		Number:   "4242424242424242",
		ExpMonth: "8",
		ExpYear:  "2023",
		CVC:      "314",
	}
	c, _ := stripefunctions.CreatePaymentMethod("cus_MH4aPLRtAMUd6o", cardParams)
	fmt.Println(c)

	/*get payment intent params from api request*/
	// paymentIntentParams := &models.StripePaymentIntent{
	// 	Amount:             1000,
	// 	PaymentMethodTypes: "card",
	// }
	// c, _ := stripefunctions.CreateStripePayIntent(paymentIntentParams)
	// fmt.Println(c)

	// c, _ := stripefunctions.CreateStripeCharge(customerParams, chargeParams)
	// fmt.Println(c)

	/*initiate creation of charge*/
	// c, _ := stripefunctions.CreateStripeCharge(customerParams, chargeParams)
	// fmt.Println(c)

	/*Get Stripe Customer*/
	//c := GetStripeCutomer("cus_MH33H53kzlxe4y")
	//fmt.Println(c)

}
