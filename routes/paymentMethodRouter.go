package routes

import (
	controller "Crunch-Garage/go-stripe/controllers"

	"github.com/gorilla/mux"
)

func PaymentMethodRouter(router *mux.Router) {
	router.HandleFunc("/api/paymentmethod/create", controller.CreatePaymentMethod).Methods("POST")
}
