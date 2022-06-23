package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Test ok")

}

//GET /sensors/:sensorId/thresholds/:thresholdId
func getSensorThreshold(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["sensor_id"]
	status := r.URL.Query().Get("status")
	customers, err := ch.service.GetAllCustomers(status)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	}
	writeResponse(w, http.StatusOK, customers)
}

//POST /sensors/:sensorId/thresholds   //Include to check to see if a threshold already exists, if it does POST request isn't allowed, and a PUT request should be recommended
func postId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["sensor_id"]
	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage("This was wrong"))
	} else {
		writeResponse(w, http.StatusOK, "customer")
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	//if this doesn't work we want the application to shutdown and show you the error message with a panic
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
