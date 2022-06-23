package controllers

import (
	"SensorProject/dtos"
	"SensorProject/service"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ITemperatureContoller interface {
	addTemperature(w http.ResponseWriter, r *http.Request)
}

type temperatureController struct {
	TemperatureService service.ITemperatureService
}

func NewTemperatureController() ITemperatureContoller {
	return temperatureController{
		TemperatureService: service.NewTemperatureService(),
	}
}

func (t temperatureController) addTemperature(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//TODO: error handling
		return
	}
	defer r.Body.Close()

	var temp dtos.AddTemperatureDto

	err = json.Unmarshal(bodyBytes, &temp)
	if err != nil {
		//TODO: error handling
		return
	}

	err = t.TemperatureService.AddTemperature(temp)
	if err != nil {
		//TODO: error handling
		return
	}
	// TODO: proper response
}
