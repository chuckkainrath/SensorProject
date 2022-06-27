package controllers

import (
	"SensorProject/dtos"
	"SensorProject/middleware"
	"SensorProject/service"
	"net/http"
)

type getThresholdController struct {
	thresholdService service.IThresholdService
}

func NewGetThresholdController(thresholdService service.IThresholdService) http.Handler {
	return &getThresholdController{
		thresholdService: thresholdService,
	}
}

//GET /sensors/:sensorId/thresholds/:thresholdId
func (th *getThresholdController) getSensorThreshold(w http.ResponseWriter, r *http.Request) {
	inputDto := **middleware.GetRequestParams(r).(**dtos.InputGetThresholdDto)

	customers, err := th.thresholdService.GetSensorThreshold(inputDto.SensorID, inputDto.ThresholdID)

	if err != nil {
		middleware.AddResultToContext(r, *err, middleware.ErrorKey)
		return
	}
	middleware.AddResultToContext(r, customers, middleware.OutputDataKey)
}

func (th *getThresholdController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	th.getSensorThreshold(w, r)
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
