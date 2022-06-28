package main

import (
	"SensorProject/app"
	event "SensorProject/events"
	"SensorProject/repository"
)

func main() {
	event.StartAddTemperatureHandler()
	repository.StartPostgresDB()
	app.StartServer()
}
