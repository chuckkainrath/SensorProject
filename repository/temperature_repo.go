package repository

import (
	"SensorProject/models"

	"github.com/shopspring/decimal"
)

type ITemperatureRepo interface {
	AddTemperatureToDb(temp *models.Temperature) error
	GetLatestTemperatures(sensorId uint, limit int) ([]decimal.Decimal, error)
}

type temperatureRepo struct{}

func NewTemperatureRepo() ITemperatureRepo {
	return temperatureRepo{}
}

func (t temperatureRepo) AddTemperatureToDb(temp *models.Temperature) error {
	result := DB().Create(temp)
	return result.Error
}

func (t temperatureRepo) GetLatestTemperatures(sensorId uint, limit int) ([]decimal.Decimal, error) {
	var temps []decimal.Decimal

	result := DB().Model(&models.Temperature{}).Select("temperatures.temperature").Where("temperatures.sensor_id = ?", sensorId).Order("temperatures.created_at DESC").Limit(limit).Find(&temps)

	if result.Error != nil {
		return nil, result.Error
	}
	return temps, nil
}
