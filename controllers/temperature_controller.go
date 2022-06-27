package controllers

import (
	"SensorProject/dtos"
	"SensorProject/events"
	"SensorProject/service"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ITemperatureController interface {
	addTemperature(w http.ResponseWriter, r *http.Request)
}

type temperatureController struct {
	TemperatureService service.ITemperatureService
	tempAddChan        chan<- uint
}

func NewTemperatureController() ITemperatureController {
	return temperatureController{
		TemperatureService: service.NewTemperatureService(),
		tempAddChan:        events.GetAddTemperatureChannel(),
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
	t.tempAddChan <- temp.SensorID

	// TODO: proper response
}
