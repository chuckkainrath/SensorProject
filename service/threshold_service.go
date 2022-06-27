package service

import (
	"SensorProject/middleware/errors"
	"SensorProject/models"
	"SensorProject/repository"
)

type IThresholdService interface {
	GetSensorThreshold(sensorId uint, thresholdId uint) (*models.Threshold, *errors.AppError)
	PostNewThreshold(sensorId uint) (*models.Threshold, *errors.AppError)
}

type ThresholdService struct {
	repo repository.ThresholdRepository
}

func (t ThresholdService) GetSensorThreshold(sensorId uint, thresholdId uint) (*models.Threshold, *errors.AppError) {
	c, err := t.repo.GetSensorThreshold(sensorId, thresholdId)
	if err != nil {
		return nil, err
	}

	//map the domain object to our dto and return it -- responsilibity of making a dto is now on the domain)

	return c, nil
}

func (t ThresholdService) PostNewThreshold(sensorId uint) (*models.Threshold, *errors.AppError) {
	c, err := t.repo.PostNewThreshold(sensorId)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func NewThresholdService(repo repository.ThresholdRepository) ThresholdService {
	return ThresholdService{repo}
}