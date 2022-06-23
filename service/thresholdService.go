package service

import (
	"SensorProject/middleware/errors"
	"SensorProject/models"
	"SensorProject/repository"
)

type IThresholdService interface {
	GetSensorThreshold(sensorId string, thresholdId string) (*models.Threshold, *errors.AppError)
}

type ThresholdService struct {
	repo repository.ThresholdRepository
}

func (t ThresholdService) GetSensorThreshold(sensorId string, thresholdId string) (*models.Threshold, *errors.AppError) {

}

func NewThresholdService(repo repository.ThresholdRepository) ThresholdService {
	return ThresholdService{repo}
}
