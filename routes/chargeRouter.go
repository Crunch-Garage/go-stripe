package routes

import (
	controller "Crunch-Garage/go-stripe/controllers"

	"github.com/gorilla/mux"
)

func ChargeRouter(router *mux.Router) {
	router.HandleFunc("/api/charge/create", controller.CreateCharge).Methods("POST")
}
