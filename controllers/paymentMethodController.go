package controller

import (
	"Crunch-Garage/go-stripe/models"
	stripefunctions "Crunch-Garage/go-stripe/stripeFunctions"
	"encoding/json"
	"net/http"
)

/*create payment method*/
func CreatePaymentMethod(w http.ResponseWriter, r *http.Request) {
	var paymentMethod models.StripeCard

	_ = json.NewDecoder(r.Body).Decode(&paymentMethod)

	//TODO: write a switch / loop to handle the payment method type i.e card e.t.c

	//TODO: Do a regex check for the card number to check if it's a valid visa or master card
	if paymentMethod.Number == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Card number is required")
		return
	}

	if paymentMethod.ExpMonth == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Card expiry month is required")
		return
	}

	if paymentMethod.ExpYear == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Card expiry year is required")
		return
	}

	if paymentMethod.CVC == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Card code is required")
		return
	}

	if paymentMethod.Type == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Payment method type is required")
		return
	}

	/*create a pament method and attaach if to a specific user (OPTIONAL) */
	/*CreatePaymentMethod(<PASS CUSTOMER ID e.g "cus_MH4aPLRtAMUd6o">, &paymentMethod)*/
	pm, pmError := stripefunctions.CreatePaymentMethod("cus_MH4aPLRtAMUd6o", &paymentMethod)

	if pmError != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(pmError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pm)
}
