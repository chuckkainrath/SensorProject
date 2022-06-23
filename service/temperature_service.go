package service

import (
	"SensorProject/dtos"
	"SensorProject/models"
	"SensorProject/repository"
)

type ITemperatureService interface {
	AddTemperature(tempDto dtos.AddTemperatureDto) error
}

type TemperatureService struct {
	TemperatureRepo repository.ITemperatureRepo
}

func (t TemperatureService) AddTemperature(tempDto dtos.AddTemperatureDto) error {

	temp := models.Temperature{
		Temperature: float64(tempDto.Temperature),
		SensorID:    tempDto.SensorID,
	}
	return t.TemperatureRepo.AddTemperatureToDb(&temp)

}
