package app

import (
	"SensorProject/controllers"
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
	thresholdRepo := repository.NewThresholdRepositoryDB(dbClient)
	tempRepo := repository.NewTemperatureRepositoryDB(dbClient)
	sensorRepo := repository.NewSensorRepositoryDB(dbClient)
	alertRepo := repository.NewThresholdAlertRepositoryDB(dbClient)
	usersRepo := repository.NewUsersRepositoryDB(dbClient)

	// Service
	thresholdService := service.NewThresholdService(thresholdRepo, tempRepo, alertRepo)
	tempService := service.NewTemperatureService(tempRepo, dateUtil)
	sensorService := service.NewSensorService(sensorRepo)
	userService := service.NewUserService(usersRepo)

	// Handlers - Threshold
	getThresholdHandler := controllers.NewGetThresholdHandler(thresholdService)
	postThresholdHandler := controllers.NewPostThresholdHandler(thresholdService)

	// Handlers - Temperature
	postTemperatureHandler := controllers.NewPostTemperatureHandler(tempService)

	// Handlers - Stats
	getReadingsHandler := controllers.NewGetReadingsHandler(tempService)
	getStatsHandler := controllers.NewGetStatsHandler(tempService)

	// Handlers - Sensor
	getAllSensorsHandler := controllers.NewGetAllSensorsHandler(sensorService)
	getSensorHandler := controllers.NewGetSensorHandler(sensorService)
	updateSensorHandler := controllers.NewUpdateSensorHandler(sensorService)

	// Handlers - User
	userLoginHandler := controllers.NewUserLoginHandler(userService)

	// Router
	router := mux.NewRouter()
	router.Use(middleware.WriteResponse)

	// User
	router.Handle("/login", middleware.BindRequestBody(userLoginHandler, &dtos.UserDto{})).Methods(http.MethodPost)

	// Temperature
	router.Handle("/sensors/temperatures", middleware.BindRequestBody(postTemperatureHandler, &dtos.AddTemperatureDto{})).Methods(http.MethodPost)

	// Auth subrouterå
	s := router.PathPrefix("/").Subrouter()
	//s.Use(auth.JwtVerify)

	// Thresholds
	s.Handle("/sensors/{sensor_id:[0-9]+}/thresholds/{threshold_id:[0-9]+}",
		middleware.BindRequestParams(getThresholdHandler, &dtos.InputGetThresholdDto{})).Methods(http.MethodGet)

	s.Handle("/sensors/thresholds",
		middleware.BindRequestBody(postThresholdHandler, &dtos.AddThresholdDto{})).Methods(http.MethodPost)

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
	s.Handle("/sensors/{sensor_id:[0-9]+}", middleware.BindRequestParams(getSensorHandler, &dtos.SensorIdDto{})).Methods(http.MethodGet)
	s.Handle("/sensors", middleware.BindRequestBody(updateSensorHandler, &dtos.UpdateSensorDto{})).Methods(http.MethodPut)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
