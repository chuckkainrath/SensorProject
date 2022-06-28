package controllers

import (
	"SensorProject/dtos"
	"SensorProject/middleware"
	"SensorProject/service"
	"net/http"
)

type getAllSensorsHandler struct {
	SensorsService service.SensorService
}

type getSensorHandler struct {
	SensorsService service.SensorService
}

type updateSensorHandler struct {
	SensorsService service.SensorService
}

func NewGetAllSensorsHandler(sensorsService service.SensorService) http.Handler {
	return &getAllSensorsHandler{
		SensorsService: sensorsService,
	}
}

func NewGetSensorHandler(sensorsService service.SensorService) http.Handler {
	return &getSensorHandler{
		SensorsService: sensorsService,
	}
}

func NewUpdateSensorHandler(sensorsService service.SensorService) http.Handler {
	return &updateSensorHandler{
		SensorsService: sensorsService,
	}
}

func (g *getAllSensorsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.GetSensors(w, r)
}

func (g *getSensorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.GetSensorById(w, r)
}

func (u *updateSensorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u.UpdateSensor(w, r)
}

func (g *getAllSensorsHandler) GetSensors(w http.ResponseWriter, r *http.Request) {
	// TODO: use tkn.UserName to get all sensors for the specified user
	// tkn := *auth.GetTokenData(r).(*models.Token)

	sensors, err := g.SensorsService.GetSensors()
	if err != nil {
		middleware.AddResultToContext(r, err, middleware.ErrorKey)
		return
	}
	middleware.AddResultToContext(r, sensors, middleware.OutputDataKey)
}

func (g *getSensorHandler) GetSensorById(w http.ResponseWriter, r *http.Request) {
	getSensorDto := **middleware.GetRequestParams(r).(**dtos.GetSensorDto)
	// TODO: use tkn.UserName to get all sensors for the specified user
	// tkn := *auth.GetTokenData(r).(*models.Token)

	sensor, err := g.SensorsService.GetSensorById(getSensorDto.SensorID)

	if err != nil {
		middleware.AddResultToContext(r, err, middleware.ErrorKey)
		return
	}
	middleware.AddResultToContext(r, sensor, middleware.OutputDataKey)
}

func (u *updateSensorHandler) UpdateSensor(w http.ResponseWriter, r *http.Request) {

}
