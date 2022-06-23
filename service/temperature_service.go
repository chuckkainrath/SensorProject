package service

import (
	"SensorProject/dtos"
	"SensorProject/models"
	"SensorProject/repository"
)

type ITemperatureService interface {
	AddTemperature(tempDto dtos.AddTemperatureDto) error
}

type temperatureService struct {
	TemperatureRepo repository.ITemperatureRepo
}

func NewTemperatureService() ITemperatureService {
	return temperatureService{
		TemperatureRepo: repository.NewTemperatureRepo(),
	}
}

func (t temperatureService) AddTemperature(tempDto dtos.AddTemperatureDto) error {

	temp := models.Temperature{
		Temperature: float64(tempDto.Temperature),
		SensorID:    tempDto.SensorID,
	}
	return t.TemperatureRepo.AddTemperatureToDb(&temp)

}
