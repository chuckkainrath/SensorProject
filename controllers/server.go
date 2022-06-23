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

	// auth routes
	s := router.PathPrefix("/auth").Subrouter()
	s.Use(auth.JwtVerify)
	// example of how we use this
	// s.HandleFunc("/user/{id}", GetUser).Methods("GET")
	log.Fatal(http.ListenAndServe("localhost:8000", router))

}
