package main

import (
	"SensorProject/controllers"
	event "SensorProject/events"
	"SensorProject/repository"
)

func main() {
	event.StartAddTemperatureHandler()
	repository.StartPostgresDB()
	controllers.StartServer()
}
