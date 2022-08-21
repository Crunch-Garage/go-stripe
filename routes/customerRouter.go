package routes

import (
	controller "Crunch-Garage/go-stripe/controllers"

	"github.com/gorilla/mux"
)

func CustomerRouter(router *mux.Router) {
	router.HandleFunc("/api/customer/create", controller.CreateStripeCustomer).Methods("POST")
}
