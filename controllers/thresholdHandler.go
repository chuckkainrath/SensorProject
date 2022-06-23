package controllers

import (
	"SensorProject/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type ThresholdHandler struct {
	service service.ThresholdService
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	//if this doesn't work we want the application to shutdown and show you the error message with a panic
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

//GET /sensors/:sensorId/thresholds/:thresholdId
func (th *ThresholdHandler) getSensorThreshold(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sensorId := vars["sensor_id"]
	thresholdId := vars["threshold_id"]
	//status := r.URL.Query().Get("status")
	customers, err := th.service.GetSensorThreshold(sensorId, thresholdId)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	}
	writeResponse(w, http.StatusOK, customers)
}

//POST /sensors/:sensorId/thresholds   //Include to check to see if a threshold already exists, if it does POST request isn't allowed, and a PUT request should be recommended
//func (th *ThresholdHandler) postId(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["sensor_id"]
// 	customer, err := th.service.GetSensorThreshold(id)
// 	if err != nil {
// 		writeResponse(w, err.Code, err.AsMessage())
// 	} else {
// 		writeResponse(w, http.StatusOK, "customer")
// 	}
// }
