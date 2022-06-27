package main

import (
	"SensorProject/controllers"
	"SensorProject/events"
	"SensorProject/repository"
)

func main() {
	events.StartAddTemperatureHandler()
	repository.StartPostgresDB()
	controllers.StartServer()
}
