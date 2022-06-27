package controllers

import (
	"SensorProject/middleware/auth"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/test", test)
	router.HandleFunc("/login", NewUserController().Login).Methods("POST")
	router.HandleFunc("/sensors", NewSensorsController().GetSensors).Methods(http.MethodGet)
	router.HandleFunc("/sensors/temperatures", NewTemperatureController().addTemperature).Methods(http.MethodPost)
	// auth routes
	s := router.PathPrefix("/auth").Subrouter()
	s.Use(auth.JwtVerify)
	// example of how we use this
	// s.HandleFunc("/user/{id}", GetUser).Methods("GET")
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
