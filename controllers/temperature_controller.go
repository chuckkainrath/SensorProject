package controllers

import (
	"SensorProject/dtos"
	"SensorProject/middleware"
	"SensorProject/service"
	"net/http"
)

type postTemperatureHandler struct {
	TemperatureService service.ITemperatureService
	tempAddChan        chan<- uint
}

func NewPostTemperatureHandler() http.Handler {
	return &postTemperatureHandler{
		TemperatureService: service.NewTemperatureService(),
		//tempAddChan:        events.GetAddTemperatureChannel(),
	}
}

func (p *postTemperatureHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.postTemperature(w, r)
}

func (p *postTemperatureHandler) postTemperature(w http.ResponseWriter, r *http.Request) {
	tempDto := **middleware.GetRequestBody(r).(**dtos.AddTemperatureDto)

	err := p.TemperatureService.AddTemperature(tempDto.SensorID, tempDto.Temperature)
	if err != nil {
		middleware.AddResultToContext(r, *err, middleware.ErrorKey)
		return
	}
	//t.tempAddChan <- temp.SensorID
}
