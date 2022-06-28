package controllers

import (
	"SensorProject/middleware"
	"SensorProject/service"
	"net/http"
)

type getAllSensorsHandler struct {
	SensorsService service.SensorsService
}

func NewGetAllSensorsHandler(sensorsService service.SensorsService) http.Handler {
	return &getAllSensorsHandler{
		SensorsService: sensorsService,
	}
}

func (g *getAllSensorsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.GetSensors(w, r)
}

// TODO: USER ID
func (g *getAllSensorsHandler) GetSensors(w http.ResponseWriter, r *http.Request) {
	sensors, err := g.SensorsService.GetSensorsService()
	if err != nil {
		middleware.AddResultToContext(r, *err, middleware.ErrorKey)
		return
	}
	middleware.AddResultToContext(r, sensors, middleware.OutputDataKey)
}
