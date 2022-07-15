package controllers

import (
	"net/http"

	"github.com/chuckkainrath/SensorProject/dtos"
	"github.com/chuckkainrath/SensorProject/middleware"
	"github.com/chuckkainrath/SensorProject/middleware/auth"
	"github.com/chuckkainrath/SensorProject/models"
	"github.com/chuckkainrath/SensorProject/service"
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

type postSensorHandler struct {
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

func NewPostSensorHandler(sensorsService service.SensorService) http.Handler {
	return &postSensorHandler{
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

func (p *postSensorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.PostSensor(w, r)
}

func (g *getAllSensorsHandler) GetSensors(w http.ResponseWriter, r *http.Request) {
	// TODO: use tkn.UserName to get all sensors for the specified user
	// tkn := *auth.GetTokenData(r).(*models.Token)
	// tkn.UserID

	sensors, err := g.SensorsService.GetSensors()
	if err != nil {
		middleware.AddResultToContext(r, err, middleware.ErrorKey)
		return
	}
	middleware.AddResultToContext(r, sensors, middleware.OutputDataKey)
}

func (g *getSensorHandler) GetSensorById(w http.ResponseWriter, r *http.Request) {
	getSensorDto := **middleware.GetRequestParams(r).(**dtos.SensorIdDto)
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
	updateSensorDto := **middleware.GetRequestBody(r).(**dtos.UpdateSensorDto)

	err := u.SensorsService.UpdateSensor(updateSensorDto.SensorID, updateSensorDto.Name, updateSensorDto.UserID)

	if err != nil {
		middleware.AddResultToContext(r, err, middleware.ErrorKey)
		return
	}

}

func (p *postSensorHandler) PostSensor(w http.ResponseWriter, r *http.Request) {
	postSensorDto := **middleware.GetRequestBody(r).(**dtos.PostSensorDto)
	// TODO: use tkn.UserName to get all sensors for the specified user

	tkn := *auth.GetTokenData(r).(*models.Token)
	err := p.SensorsService.PostSensor(postSensorDto.Name, tkn.UserID)

	if err != nil {
		middleware.AddResultToContext(r, err, middleware.ErrorKey)
		return
	}
}
