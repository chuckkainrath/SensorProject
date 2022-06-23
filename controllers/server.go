package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/test", test)

	log.Fatal(http.ListenAndServe("localhost:8000", router))

}
