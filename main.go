package main

import (
	event "SensorProject/events"
	"SensorProject/repository"
)

func main() {
	event.StartAddTemperatureHandler()
	repository.StartPostgresDB()
	StartServer()
}
