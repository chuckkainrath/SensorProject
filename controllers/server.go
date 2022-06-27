package controllers

import (
	"SensorProject/repository"
	"SensorProject/service"
	"SensorProject/util"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/test", test)

	db := repository.DB()
	tempRepo := repository.NewTemperatureRepo(db)
	dateUtil := util.NewDateChecker()
	tempService := service.NewTemperatureService(tempRepo, dateUtil)
	tempController := NewTemperatureController(tempService)

	router.HandleFunc("/sensors/{sensor_id:[0-9]+}/stats/readings", tempController.GetPerMinuteReadings).
		Methods(http.MethodGet).
		Queries("from", "{from}").
		Queries("to", "{to}")
	router.HandleFunc("/sensors/{sensor_id:[0-9]+}/stats/minmaxaverage", tempController.GetMinMaxAverageStats).
		Methods(http.MethodGet).
		Queries("from", "{from}").
		Queries("to", "{to}")

	log.Fatal(http.ListenAndServe("localhost:8000", router))

}
