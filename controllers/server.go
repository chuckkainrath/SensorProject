package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/test", test)
	router.HandleFunc("/sensors/{sensor_id:[0-9]+}/thresholds", postId).Methods(http.MethodPost)
	router.HandleFunc("/sensors/{sensor_id:[0-9]+}/thresholds/{threshold_id: [0-9]+}", getSensorThreshold).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}
