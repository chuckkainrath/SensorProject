package controllers

import (
	"net/http"

	"github.com/chuckkainrath/SensorProject/dtos"
	"github.com/chuckkainrath/SensorProject/middleware"
	"github.com/chuckkainrath/SensorProject/middleware/auth"
	"github.com/chuckkainrath/SensorProject/models"
	"github.com/chuckkainrath/SensorProject/service"
)

type getReadings struct {
	TemperatureService service.TemperatureService
}

type getStats struct {
	TemperatureService service.TemperatureService
}

func NewGetReadingsHandler(temperatureService service.TemperatureService) http.Handler {
	return &getReadings{
		TemperatureService: temperatureService,
	}
}

func NewGetStatsHandler(temperatureService service.TemperatureService) http.Handler {
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
	userTkn := *auth.GetTokenData(r).(*models.Token)

	res, err := g.TemperatureService.GetPerMinuteReading(statsDto.SensorID, statsDto.From, statsDto.To, userTkn.UserID)
	if err != nil {
		middleware.AddResultToContext(r, err, middleware.ErrorKey)
		return
	}
	middleware.AddResultToContext(r, res, middleware.OutputDataKey)
}

func (g *getStats) GetMinMaxAverageStats(w http.ResponseWriter, r *http.Request) {
	statsDto := **middleware.GetRequestParams(r).(**dtos.InputStatsDto)
	userTkn := *auth.GetTokenData(r).(*models.Token)

	res, err := g.TemperatureService.GetMinMaxAverageStats(statsDto.SensorID, statsDto.From, statsDto.To, userTkn.UserID)
	if err != nil {
		middleware.AddResultToContext(r, err, middleware.ErrorKey)
		return
	}
	middleware.AddResultToContext(r, res, middleware.OutputDataKey)
}
