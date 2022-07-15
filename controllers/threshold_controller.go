package controllers

import (
	"net/http"

	"github.com/chuckkainrath/SensorProject/dtos"
	"github.com/chuckkainrath/SensorProject/middleware"
	"github.com/chuckkainrath/SensorProject/service"
)

type getThresholdHandler struct {
	thresholdService service.ThresholdService
}

type postThresholdHandler struct {
	thresholdService    service.ThresholdService
	updateThresholdChan chan<- dtos.ThresholdEventDto
}

type deleteThresholdHandler struct {
	thresholdService    service.ThresholdService
	updateThresholdChan chan<- dtos.ThresholdEventDto
}

func NewGetThresholdHandler(thresholdService service.ThresholdService) http.Handler {
	return &getThresholdHandler{
		thresholdService: thresholdService,
	}
}

func NewPostThresholdHandler(thresholdService service.ThresholdService,
	updateThresholdChan chan<- dtos.ThresholdEventDto) http.Handler {
	return &postThresholdHandler{
		thresholdService:    thresholdService,
		updateThresholdChan: updateThresholdChan,
	}
}

func NewDeleteThresholdHandler(thresholdService service.ThresholdService,
	updateThresholdChan chan<- dtos.ThresholdEventDto) http.Handler {
	return &deleteThresholdHandler{
		thresholdService:    thresholdService,
		updateThresholdChan: updateThresholdChan,
	}
}

// func NewPutThresholdHandler(thresholdService service.ThresholdService) http.Handler {
// 	return &putThresholdHandler{
// 		thresholdService: thresholdService,
// 	}
// }

//GET /sensors/:sensorId/thresholds/:thresholdId
func (th *getThresholdHandler) getSensorThreshold(w http.ResponseWriter, r *http.Request) {
	inputDto := **middleware.GetRequestParams(r).(**dtos.InputGetThresholdDto)
	threshold, err := th.thresholdService.GetSensorThreshold(inputDto.SensorID)

	if err != nil {
		middleware.AddResultToContext(r, err, middleware.ErrorKey)
		return
	}
	middleware.AddResultToContext(r, threshold, middleware.OutputDataKey)
}

//POST /sensors/thresholds   //Include to check to see if a threshold already exists, if it does POST request isn't allowed, and a PUT request should be recommended
func (th *postThresholdHandler) postSensorThreshold(w http.ResponseWriter, r *http.Request) {
	inputDto := **middleware.GetRequestBody(r).(**dtos.AddThresholdDto)

	err := th.thresholdService.UpsertNewThreshold(inputDto.SensorID, inputDto.Temperature)
	if err != nil {
		middleware.AddResultToContext(r, err, middleware.ErrorKey)
		return
	}

	th.updateThresholdChan <- dtos.ThresholdEventDto{
		SensorID:    inputDto.SensorID,
		Temperature: &inputDto.Temperature,
	}
}

func (th *deleteThresholdHandler) deleteSensorThreshold(w http.ResponseWriter, r *http.Request) {
	inputDto := **middleware.GetRequestParams(r).(**dtos.InputGetThresholdDto)
	err := th.thresholdService.DeleteSensorThreshold(inputDto.SensorID)

	if err != nil {
		middleware.AddResultToContext(r, err, middleware.ErrorKey)
		return
	}

	th.updateThresholdChan <- dtos.ThresholdEventDto{
		SensorID:    inputDto.SensorID,
		Temperature: nil,
	}
}

func (th *getThresholdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	th.getSensorThreshold(w, r)
}

func (th *postThresholdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	th.postSensorThreshold(w, r)
}

func (th *deleteThresholdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	th.deleteSensorThreshold(w, r)
}
