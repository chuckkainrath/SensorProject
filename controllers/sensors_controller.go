package controllers

import (
	"SensorProject/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ISensorsController interface {
	GetSensors(w http.ResponseWriter, r *http.Request)
	GetSensorId(w http.ResponseWriter, r *http.Request)
	UpdateSensor(w http.ResponseWriter, r *http.Request)
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

func (s sensorsController) GetSensorId(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	stringId := vars["id"]
	sensorId, _ := strconv.ParseInt(stringId, 10, 0)

	SensorsResponse, err := s.sensorsService.GetSensorIdService(int(sensorId))

	if err != nil {
		//TODO:
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	jsonResp, err := json.Marshal(SensorsResponse)
	if err != nil {
		panic(err)
	}
	w.Write(jsonResp)

}

func (s sensorsController) UpdateSensor(w http.ResponseWriter, r *http.Request) {

}
