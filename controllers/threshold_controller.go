package controllers

import (
	"SensorProject/dtos"
	"SensorProject/middleware"
	"SensorProject/service"
	"net/http"
)

type getThresholdHandler struct {
	thresholdService service.ThresholdService
}

type postThresholdHandler struct {
	thresholdService service.ThresholdService
}

func NewGetThresholdHandler(thresholdService service.ThresholdService) http.Handler {
	return &getThresholdHandler{
		thresholdService: thresholdService,
	}
}

func NewPostThresholdHandler(thresholdService service.ThresholdService) http.Handler {
	return &postThresholdHandler{
		thresholdService: thresholdService,
	}
}

//GET /sensors/:sensorId/thresholds/:thresholdId
func (th *getThresholdHandler) getSensorThreshold(w http.ResponseWriter, r *http.Request) {
	inputDto := **middleware.GetRequestParams(r).(**dtos.InputGetThresholdDto)

	customers, err := th.thresholdService.GetSensorThreshold(inputDto.SensorID, inputDto.ThresholdID)

	if err != nil {
		middleware.AddResultToContext(r, err, middleware.ErrorKey)
		return
	}
	middleware.AddResultToContext(r, customers, middleware.OutputDataKey)
}

//POST /sensors/:sensorId/thresholds   //Include to check to see if a threshold already exists, if it does POST request isn't allowed, and a PUT request should be recommended
func (th *postThresholdHandler) postSensorThreshold(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// id, _ := strconv.Atoi(vars["sensor_id"])
	// customer, err := th.service.PostNewThreshold(id)
	// if err != nil {
	// 	writeResponse(w, err.Code, err.AsMessage())
	// } else {
	// 	writeResponse(w, http.StatusOK, customer)
	// }
}

func (th *getThresholdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	th.getSensorThreshold(w, r)
}

func (th *postThresholdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	th.postSensorThreshold(w, r)
}
