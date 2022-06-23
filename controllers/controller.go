package controllers

// POST /sensors/:sensorId/thresholds //Include to check to see if a threshold already exists, if it does POST request isn't allowed, and a PUT request should be recommended
// GET /sensors/:sensorId/thresholds/:thresholdId
// PUT /sensors/:sensorId/thresholds/:thresholdId
// DELETE /sensors/:sensorId/thresholds/:thresholdId
import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	router.HandleFunc("/sensors/{sensor_id:[0-9]+}/thresholds", postId).Methods(http.MethodPost)
}
func postId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["sensor_id"]
	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
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
