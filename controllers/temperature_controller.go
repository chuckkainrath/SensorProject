package controllers

import (
	"SensorProject/service"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
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

func (g *getStats) GetMinMaxAverageStats(w http.ResponseWriter, r *http.Request) {
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
