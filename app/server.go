package app

import (
	"log"
	"net/http"

	"github.com/chuckkainrath/SensorProject/controllers"
	"github.com/chuckkainrath/SensorProject/dtos"
	event "github.com/chuckkainrath/SensorProject/events"
	"github.com/chuckkainrath/SensorProject/middleware"
	"github.com/chuckkainrath/SensorProject/middleware/auth"
	"github.com/chuckkainrath/SensorProject/repository"
	"github.com/chuckkainrath/SensorProject/service"
	"github.com/chuckkainrath/SensorProject/util"
	"github.com/gorilla/mux"
)

func StartServer() {
	dbClient := repository.DB()

	// Channels
	tempAddChan := event.GetAddTemperatureChannel()
	thresholdUpdateChan := event.GetUpdateThresholdChannel()

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
	// postThresholdHandler := controllers.NewPostThresholdHandler(thresholdService)
	deleteThresholdHandler := controllers.NewDeleteThresholdHandler(thresholdService, thresholdUpdateChan)
	postThresholdHandler := controllers.NewPostThresholdHandler(thresholdService, thresholdUpdateChan)

	// Handlers - Temperature
	postTemperatureHandler := controllers.NewPostTemperatureHandler(tempService, tempAddChan)

	// Handlers - Stats
	getReadingsHandler := controllers.NewGetReadingsHandler(tempService)
	getStatsHandler := controllers.NewGetStatsHandler(tempService)

	// Handlers - Sensor
	getAllSensorsHandler := controllers.NewGetAllSensorsHandler(sensorService)
	getSensorHandler := controllers.NewGetSensorHandler(sensorService)
	updateSensorHandler := controllers.NewUpdateSensorHandler(sensorService)
	postSensorHandler := controllers.NewPostSensorHandler((sensorService))

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
	s.Use(auth.JwtVerify)

	// Thresholds
	s.Handle("/sensors/{sensor_id:[0-9]+}/thresholds",
		middleware.BindRequestParams(getThresholdHandler, &dtos.InputGetThresholdDto{})).Methods(http.MethodGet)

	s.Handle("/sensors/thresholds",
		middleware.BindRequestBody(postThresholdHandler, &dtos.AddThresholdDto{})).Methods(http.MethodPost, http.MethodPut)

	s.Handle("/sensors/{sensor_id:[0-9]+}/thresholds",
		middleware.BindRequestBody(deleteThresholdHandler, &dtos.InputGetThresholdDto{})).Methods(http.MethodDelete)

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
	s.Handle("/sensors/{sensor_id:[0-9]+}", middleware.BindRequestBody(updateSensorHandler, &dtos.UpdateSensorDto{})).Methods(http.MethodPut)
	s.Handle("/sensors", middleware.BindRequestBody(postSensorHandler, &dtos.PostSensorDto{})).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
