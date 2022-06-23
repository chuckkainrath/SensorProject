package controllers

import (
	"SensorProject/dtos"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func addTemperature(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//TODO: error handling
		return
	}
	defer r.Body.Close()

	var temp dtos.AddTemperatureDto

	err = json.Unmarshal(bodyBytes, &temp)
	if err != nil {
		//TODO: error handling
		return
	}

	

}
