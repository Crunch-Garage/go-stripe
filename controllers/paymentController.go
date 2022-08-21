package controller

import (
	"encoding/json"
	"net/http"
)

func MakePayment(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode("write response")
}
