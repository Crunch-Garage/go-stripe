package controller

import (
	"Crunch-Garage/go-stripe/models"
	stripefunctions "Crunch-Garage/go-stripe/stripeFunctions"
	"encoding/json"
	"net/http"
)

/*create stripe customer*/
func CreateStripeCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.StripeCustomer

	_ = json.NewDecoder(r.Body).Decode(&customer)

	/*customer parameters are optional*/
	newCustomer, err := stripefunctions.CreateStripeCustomer(&customer)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newCustomer)
}
