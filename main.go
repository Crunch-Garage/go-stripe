package main

import (
	"Crunch-Garage/go-stripe/connections"
	"Crunch-Garage/go-stripe/routes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	/*initiate stripe*/
	connections.GlobalStripeConnection()

	/*Listen and serve incoming requests*/
	HandleApiRequests()

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
	// cardParams := &models.StripeCard{
	// 	Number:   "4242424242424242",
	// 	ExpMonth: "8",
	// 	ExpYear:  "2023",
	// 	CVC:      "314",
	// }
	// c, _ := stripefunctions.CreatePaymentMethod("cus_MH4aPLRtAMUd6o", cardParams)
	// fmt.Println(c)

	/*get payment intent params from api request*/
	// paymentIntentParams := &models.StripePaymentIntent{
	// 	Amount:             1000,
	// 	PaymentMethodTypes: "card",

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

/*Handle Api Requests*/
func HandleApiRequests() {

	router := mux.NewRouter()

	routes.PaymentRouter(router)

	srv := &http.Server{
		Addr: ":8080", // "127.0.0.1:8000"  "0.0.0.0:8000"
		/*Good practice to set timeouts to avoid Slowloris attacks.*/
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, // Pass our instance of gorilla/mux in.
	}

	log.Fatal(srv.ListenAndServe()) //log.Fatal(http.ListenAndServe(":8080", router))

}
