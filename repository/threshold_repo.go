package repository

import (
	"SensorProject/models"

	"github.com/shopspring/decimal"
)

type IThresholdRepo interface {
	GetThresholdTemperature(sensorId uint) (*decimal.Decimal, error)
}

type thresholdRepo struct{}

func NewThresholdRepo() IThresholdRepo {
	return thresholdRepo{}
}

func (t thresholdRepo) GetThresholdTemperature(sensorId uint) (*decimal.Decimal, error) {
	var threshold models.Threshold
	result := DB().Model(&models.Threshold{}).Select("thresholds.temperature").Where("thresholds.sensor_id = ?", sensorId).First(&threshold)
	if result.Error != nil {
		return nil, result.Error
	}
	return &(threshold.Temperature), nil
}
