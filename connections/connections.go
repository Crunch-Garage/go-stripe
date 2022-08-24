package connections

import (
	"Crunch-Garage/go-stripe/config"

	"github.com/stripe/stripe-go"
)

func GlobalStripeConnection() {

	/*GLOBALY SET API KEY*/
	/*initialize stripe client globally if using one stripe account*/
	stripe.Key = config.EnvStripeApiKey()

	/*SET API KEY PER REQUEST*/
	/* 1- if using multiple stripe accounts say one in USA and another in Europe
	   2- set an api key each time when making a new api per  request
	   3- if building a platform for multiple vendors i.e UBER, LYFT   pass the api key in requests
	*/
	//sc := &client.API{}
	//sc.Init(config.EnvStripeApiKey(), nil)
	//c, _ := sc.Customers.Get("cus_MH2lQHxvjab6TA", nil)
}
