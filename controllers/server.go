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

	dateUtil := util.NewDateChecker()

	thresholdRepository := repository.NewThresholdRepositoryDB(dbClient)
	tempRepo := repository.NewTemperatureRepositoryDB(dbClient)

	thresholdService := service.NewThresholdService(thresholdRepository)
	tempService := service.NewTemperatureService(tempRepo, dateUtil)

	getThresholdHandler := NewGetThresholdController(thresholdService)

	getReadingsHandler := NewGetReadingsHandler(tempService)
	getStatsHandler := NewGetStatsHandler(tempService)

	router := mux.NewRouter()
	router.HandleFunc("/test", test)
	//router.HandleFunc("/sensors/{sensor_id:[0-9]+}/thresholds", th.postId).Methods(http.MethodPost)
	//router.HandleFunc("/sensors/{sensor_id:[0-9]+}/thresholds/{threshold_id: [0-9]+}", thresholdController.GetSensorThreshold).Methods(http.MethodGet)
	router.Handle("/sensors/{sensor_id:[0-9]+}/thresholds/{threshold_id: [0-9]+}",
		middleware.BindParams(middleware.WriteResponse(getThresholdHandler), &dtos.InputGetThresholdDto{}))

	router.Handle("/sensors/{sensor_id:[0-9]+}/stats/readings",
		middleware.BindParams(middleware.WriteResponse(getThresholdHandler), &dtos.InputStatsDto{})).
		Methods(http.MethodGet).
		Queries("from", "{from}").
		Queries("to", "{to}")
	router.Handle("/sensors/{sensor_id:[0-9]+}/stats/minmaxaverage",
		middleware.BindParams(middleware.WriteResponse(getThresholdHandler), &dtos.InputStatsDto{})).
		Methods(http.MethodGet).
		Queries("from", "{from}").
		Queries("to", "{to}")

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
