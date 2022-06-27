package controllers

import (
	"SensorProject/service"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type ITemperatureContoller interface {
	GetPerMinuteReadings(w http.ResponseWriter, r *http.Request)
	GetMinMaxAverageStats(w http.ResponseWriter, r *http.Request)
}

type temperatureController struct {
	TemperatureService service.ITemperatureService
}

func NewTemperatureController(temperatureService service.ITemperatureService) ITemperatureContoller {
	return temperatureController{
		TemperatureService: temperatureService,
	}
}

func (t temperatureController) GetPerMinuteReadings(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sensorId, err := strconv.Atoi(vars["sensor_id"])

	layout := "2006-01-02T15:04:05"
	from, err1 := time.Parse(layout, vars["from"])
	to, err2 := time.Parse(layout, vars["to"])

	if err != nil || err1 != nil || err2 != nil {
		// TODO: handle error response
		return
	}

	res, err := t.TemperatureService.GetPerMinuteReading(uint(sensorId), from, to)
	if err != nil {
		// TODO: handle error response
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (t temperatureController) GetMinMaxAverageStats(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sensorId, err := strconv.Atoi(vars["sensor_id"])

	layout := "2006-01-02T15:04:05"

	from, err1 := time.Parse(layout, vars["from"])
	to, err2 := time.Parse(layout, vars["to"])

	if err != nil || err1 != nil || err2 != nil {
		// TODO: handle error response
		return
	}

	res, err := t.TemperatureService.GetMinMaxAverageStats(uint(sensorId), from, to)
	if err != nil {
		// TODO: handle error response
		return
	}

	json.NewEncoder(w).Encode(res)
}
