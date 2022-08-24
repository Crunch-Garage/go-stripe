package controller

import (
	"Crunch-Garage/go-stripe/models"
	stripefunctions "Crunch-Garage/go-stripe/stripeFunctions"
	"encoding/json"
	"net/http"
)

/*create stripe charege controller*/
func CreateCharge(w http.ResponseWriter, r *http.Request) {
	var charge models.StripeCharge

	_ = json.NewDecoder(r.Body).Decode(&charge)

	if charge.Token == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Charge token is required")
		return
	}

	if charge.Amount == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Inavlid Amount")
		return
	}

	/*initiate creation of charge*/
	newCharge, chargeError := stripefunctions.CreateStripeCharge(&charge)

	if chargeError != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(chargeError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newCharge)
}
