package controllers

import (
	"SensorProject/dtos"
	"SensorProject/service"
	"encoding/json"
	"net/http"
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
	requestDto := &dtos.TemperatureStatsInputDto{}
	err := json.NewDecoder(r.Body).Decode(requestDto)
	if err != nil {
		// TODO: handle error response
		return
	}

	res, err := t.TemperatureService.GetPerMinuteReading(requestDto.SensorId, requestDto.From, requestDto.To)
	if err != nil {
		// TODO: handle error response
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (t temperatureController) GetMinMaxAverageStats(w http.ResponseWriter, r *http.Request) {
	requestDto := &dtos.TemperatureStatsInputDto{}
	err := json.NewDecoder(r.Body).Decode(requestDto)
	if err != nil {
		// TODO: handle error response
		return
	}

	res, err := t.TemperatureService.GetMinMaxAverageStats(requestDto.SensorId, requestDto.From, requestDto.To)
	if err != nil {
		// TODO: handle error response
		return
	}

	json.NewEncoder(w).Encode(res)
}
