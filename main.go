package main

import (
	"github.com/chuckkainrath/SensorProject/app"
	event "github.com/chuckkainrath/SensorProject/events"
	"github.com/chuckkainrath/SensorProject/repository"
)

func main() {
	event.StartAddTemperatureHandler()
	repository.StartPostgresDB()
	app.StartServer()
}
