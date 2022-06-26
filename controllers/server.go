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

	router.HandleFunc("/stats/readings", tempController.GetPerMinuteReadings).Methods(http.MethodPost)
	router.HandleFunc("/stats/minmaxaverage", tempController.GetMinMaxAverageStats).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:8000", router))

}
