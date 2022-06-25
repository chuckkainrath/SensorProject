package repository

import (
	"SensorProject/dtos"
	"time"
)

type TemperatureRepo interface {
	GetPerMinuteReading(sensorId uint, to, from time.Time) (*[]dtos.TemperatureDto, error)
}

type 

func NewTemperatureRepo(temperatureRepo TemperatureRepo) TemperatureRepo {

}