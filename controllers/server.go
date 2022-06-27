package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/test", test)
	router.HandleFunc("/sensors", NewSensorsController().GetSensors).Methods(http.MethodGet)
	router.HandleFunc("/sensors/{id:[0-9]+}", NewSensorsController().GetSensorId).Methods(http.MethodGet)
	router.HandleFunc("/sensors/{id:[0-9]+}", NewSensorsController().UpdateSensor).Methods(http.MethodPut)
	router.HandleFunc("/sensors/temperatures", NewTemperatureController().addTemperature).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:8000", router))

}
