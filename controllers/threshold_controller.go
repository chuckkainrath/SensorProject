package controllers

import (
	"SensorProject/dtos"
	"SensorProject/middleware"
	"SensorProject/middleware/auth"
	"SensorProject/models"
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
	user := *auth.GetTokenData(r).(*models.Token)

	threshold, err := th.thresholdService.GetSensorThreshold(inputDto.SensorID, user.UserID)

	if err != nil {
		middleware.AddResultToContext(r, err, middleware.ErrorKey)
		return
	}
	middleware.AddResultToContext(r, threshold, middleware.OutputDataKey)
}

//POST /sensors/thresholds   //Include to check to see if a threshold already exists, if it does POST request isn't allowed, and a PUT request should be recommended
func (th *postThresholdHandler) postSensorThreshold(w http.ResponseWriter, r *http.Request) {
	inputDto := **middleware.GetRequestBody(r).(**dtos.AddThresholdDto)
	//user := *auth.GetTokenData(r).(*models.Token)
	//TODO: DUSTIN, why is this not working? Try removing newThreshold?
	err := th.thresholdService.UpsertNewThreshold(inputDto.SensorID, inputDto.Temperature)
	if err != nil {
		middleware.AddResultToContext(r, err, middleware.ErrorKey)
		return
	}
}

func (th *getThresholdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	th.getSensorThreshold(w, r)
}

func (th *postThresholdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	th.postSensorThreshold(w, r)
}
