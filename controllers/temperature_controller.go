package controllers

import (
	"SensorProject/dtos"
	"SensorProject/middleware"
	"SensorProject/service"
	"net/http"
)

type getReadings struct {
	TemperatureService service.ITemperatureService
}

type getStats struct {
	TemperatureService service.ITemperatureService
}

func NewGetReadingsHandler(temperatureService service.ITemperatureService) http.Handler {
	return &getReadings{
		TemperatureService: temperatureService,
	}
}

func NewGetStatsHandler(temperatureService service.ITemperatureService) http.Handler {
	return &getStats{
		TemperatureService: temperatureService,
	}
}

func (g *getReadings) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.GetPerMinuteReadings(w, r)
}

func (g *getStats) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.GetMinMaxAverageStats(w, r)
}

func (g *getReadings) GetPerMinuteReadings(w http.ResponseWriter, r *http.Request) {
	statsDto := **middleware.GetRequestParams(r).(**dtos.InputStatsDto)

	res, err := g.TemperatureService.GetPerMinuteReading(statsDto.SensorID, statsDto.From, statsDto.To)
	if err != nil {
		middleware.AddResultToContext(r, *err, middleware.ErrorKey)
		return
	}
	middleware.AddResultToContext(r, res, middleware.OutputDataKey)
}

func (g *getStats) GetMinMaxAverageStats(w http.ResponseWriter, r *http.Request) {
	statsDto := **middleware.GetRequestParams(r).(**dtos.InputStatsDto)

	res, err := g.TemperatureService.GetMinMaxAverageStats(statsDto.SensorID, statsDto.From, statsDto.To)
	if err != nil {
		middleware.AddResultToContext(r, *err, middleware.ErrorKey)
		return
	}
	middleware.AddResultToContext(r, res, middleware.OutputDataKey)
}
