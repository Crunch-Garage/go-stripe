package routes

import (
	controller "Crunch-Garage/go-stripe/controllers"

	"github.com/gorilla/mux"
)

func PaymentRouter(router *mux.Router) {
	router.HandleFunc("/api/payment", controller.MakePayment).Methods("POST")
}
