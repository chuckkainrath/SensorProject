package repository

import "SensorProject/models"

type IThresholdAlertRepo interface {
	AddThresholdAlert(alert *models.ThresholdAlert) error
}

type thresholdAlertRepo struct{}

func NewThresholdAlertRepo() IThresholdAlertRepo {
	return thresholdAlertRepo{}
}

func (t thresholdAlertRepo) AddThresholdAlert(alert *models.ThresholdAlert) error {
	result := DB().Create(alert)
	return result.Error
}
