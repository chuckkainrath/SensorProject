package controllers

import (
	"SensorProject/repository"
	"SensorProject/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	dbClient := repository.DB()
	thresholdRepository := repository.NewThresholdRepositoryDB(dbClient)
	thresholdService := service.NewThresholdService(thresholdRepository)
	thresholdController := NewThresholdController(thresholdService)

	router := mux.NewRouter()
	router.HandleFunc("/test", test)
	//router.HandleFunc("/sensors/{sensor_id:[0-9]+}/thresholds", th.postId).Methods(http.MethodPost)
	router.HandleFunc("/sensors/{sensor_id:[0-9]+}/thresholds/{threshold_id: [0-9]+}", thresholdController.GetSensorThreshold).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}
