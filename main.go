package main

import (
	"SensorProject/controllers"
	"SensorProject/repository"
)

func main() {

	repository.StartPostgresDB()
	controllers.StartServer()

}
