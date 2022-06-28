package controllers

import (
	"SensorProject/dtos"
	"SensorProject/middleware"
	"SensorProject/middleware/auth"
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
	thresholdRepo := repository.NewThresholdRepositoryDB(dbClient)
	tempRepo := repository.NewTemperatureRepositoryDB(dbClient)
	sensorRepo := repository.NewSensorsRepositoryDB(dbClient)
	alertRepo := repository.NewThresholdAlertRepositoryDB(dbClient)
	usersRepo := repository.NewUsersRepositoryDB(dbClient)

	// Service
	thresholdService := service.NewThresholdService(thresholdRepo, tempRepo, alertRepo)
	tempService := service.NewTemperatureService(tempRepo, dateUtil)
	sensorService := service.NewSensorsService(sensorRepo)
	userService := service.NewUserService(usersRepo)

	// Handlers - Threshold
	getThresholdHandler := NewGetThresholdHandler(thresholdService)
	postThresholdHandler := NewPostThresholdHandler(thresholdService)

	// Handlers - Temperature
	postTemperatureHandler := NewPostTemperatureHandler(tempService)

	// Handlers - Stats
	getReadingsHandler := NewGetReadingsHandler(tempService)
	getStatsHandler := NewGetStatsHandler(tempService)

	// Handlers - Sensor
	getAllSensorsHandler := NewGetAllSensorsHandler(sensorService)

	// Handlers - User
	userLoginHandler := NewUserLoginHandler(userService)

	// Router
	router := mux.NewRouter()
	router.Use(middleware.WriteResponse)

	// User
	router.Handle("/login", middleware.BindRequestBody(userLoginHandler, dtos.UserDto{})).Methods(http.MethodPost)

	// Temperature
	router.Handle("/sensors/temperatures", middleware.BindRequestBody(postTemperatureHandler, dtos.AddTemperatureDto{})).Methods(http.MethodPost)

	// Auth subrouter
	s := router.PathPrefix("/").Subrouter()
	s.Use(auth.JwtVerify)

	// Thresholds
	s.Handle("/sensors/{sensor_id:[0-9]+}/thresholds/{threshold_id: [0-9]+}",
		middleware.BindRequestParams(getThresholdHandler, &dtos.InputGetThresholdDto{})).Methods(http.MethodGet)

	s.Handle("/sensors/{sensor_id:[0-9]+}/thresholds",
		middleware.BindRequestBody(postThresholdHandler, dtos.AddThresholdDto{})).Methods(http.MethodPost)

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
	s.Handle("/sensors", getAllSensorsHandler).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
