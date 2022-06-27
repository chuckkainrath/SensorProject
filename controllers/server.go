package controllers

import (
	"SensorProject/dtos"
	"SensorProject/middleware"
	"SensorProject/repository"
	"SensorProject/service"
	"SensorProject/util"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	dbClient := repository.DB()

	// Util
	dateUtil := util.NewDateChecker()

	// Repo
	thresholdRepository := repository.NewThresholdRepositoryDB(dbClient)
	tempRepo := repository.NewTemperatureRepositoryDB(dbClient)

	// Service
	thresholdService := service.NewThresholdService(thresholdRepository)
	tempService := service.NewTemperatureService(tempRepo, dateUtil)

	// Handlers
	getThresholdHandler := NewGetThresholdController(thresholdService)

	getReadingsHandler := NewGetReadingsHandler(tempService)
	getStatsHandler := NewGetStatsHandler(tempService)

	// Router
	router := mux.NewRouter()
	router.HandleFunc("/test", test)

	// Thresholds
	router.Handle("/sensors/{sensor_id:[0-9]+}/thresholds/{threshold_id: [0-9]+}",
		middleware.BindParams(middleware.WriteResponse(getThresholdHandler), &dtos.InputGetThresholdDto{}))

	// Stats
	router.Handle("/sensors/{sensor_id:[0-9]+}/stats/readings",
		middleware.BindParams(middleware.WriteResponse(getReadingsHandler), &dtos.InputStatsDto{})).
		Methods(http.MethodGet).
		Queries("from", "{from}").
		Queries("to", "{to}")
	router.Handle("/sensors/{sensor_id:[0-9]+}/stats/minmaxaverage",
		middleware.BindParams(middleware.WriteResponse(getStatsHandler), &dtos.InputStatsDto{})).
		Methods(http.MethodGet).
		Queries("from", "{from}").
		Queries("to", "{to}")

	// Sensors

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
