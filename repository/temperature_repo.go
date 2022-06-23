package repository

import "SensorProject/models"

type ITemperatureRepo interface {
	AddTemperatureToDb(temp *models.Temperature) error
}


type TemperatureRepo struct{}

func(t TemperatureRepo)AddTemperatureToDb(temp *models.Temperature) error{

	result:=DB().Create(temp)
	return result.Error
}