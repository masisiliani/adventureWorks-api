package controller

import (
	"adventureWorks-api/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetCurrentMatchs return matchs will occurs today
func GetCustomerByID(w http.ResponseWriter, r *http.Request) {

	idCustomer, err := strconv.Atoi(mux.Vars(r)["id"])
	if err == nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	objCustomer, err := service.GetCustomerByID(idCustomer)
	if err == nil {

		w.Header().Set(`Content-Type`, `application/json`)
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(objCustomer)

	} else {

		w.WriteHeader(http.StatusBadRequest)

	}
}
