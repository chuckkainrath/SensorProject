package repository

import "SensorProject/models"

type ITemperatureRepo interface {
	AddTemperatureToDb(temp *models.Temperature) error
}

type temperatureRepo struct{}

func NewTemperatureRepo() ITemperatureRepo {
	return temperatureRepo{}
}

func (t temperatureRepo) AddTemperatureToDb(temp *models.Temperature) error {
	result := DB().Create(temp)
	return result.Error
}
