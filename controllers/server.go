package controllers

import (
	"SensorProject/dtos"
	"SensorProject/middleware"
	"SensorProject/repository"
	"SensorProject/service"
	"SensorProject/util"
	"SensorProject/middleware/auth"
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
	router.Use(middleware.WriteResponse) 

	// User
	router.HandleFunc("/login", NewUserController().Login).Methods("POST")

	// Temperature
	router.HandleFunc("/sensors/temperatures", NewTemperatureController().addTemperature).Methods(http.MethodPost)

	// Auth subrouter
	s := router.PathPrefix("/").Subrouter()
	s.Use(auth.JwtVerify)

	// Thresholds
	s.Handle("/sensors/{sensor_id:[0-9]+}/thresholds/{threshold_id: [0-9]+}",
		middleware.BindRequestParams(getThresholdHandler, &dtos.InputGetThresholdDto{}))

	// Stats
	s.Handle("/sensors/{sensor_id:[0-9]+}/stats/readings",
		middleware.BindRequestParams(getReadingsHandler, &dtos.InputStatsDto{})).
		Methods(http.MethodGet).
		Queries("from", "{from}").
		Queries("to", "{to}")
	s.Handle("/sensors/{sensor_id:[0-9]+}/stats/minmaxaverage",
		middleware.BindRequestParams(getStatsHandler, &dtos.InputStatsDto{})).
		Methods(http.MethodGet).
		Queries("from", "{from}").
		Queries("to", "{to}")


	// Sensors
	s.HandleFunc("/sensors", NewSensorsController().GetSensors).Methods(http.MethodGet)
	
	// auth routes
	

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
