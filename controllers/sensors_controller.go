package controllers

import (
	"SensorProject/service"
	"encoding/json"
	"net/http"
)

type ISensorsController interface {
	GetSensors(w http.ResponseWriter, r *http.Request)
	
}

type sensorsController struct {
	sensorsService service.ISensorsService
}

func NewSensorsController() ISensorsController {
	return sensorsController{
		sensorsService: service.NewSensorsService(),
	}
}

func (s sensorsController) GetSensors(w http.ResponseWriter, r *http.Request) {

	SensorsResponse := s.sensorsService.GetSensorsService()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	jsonResp, err := json.Marshal(SensorsResponse)
	if err != nil {
		panic(err)
	}
	w.Write(jsonResp)

}
