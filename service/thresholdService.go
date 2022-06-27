package service

import (
	"SensorProject/middleware/errors"
	"SensorProject/models"
	"SensorProject/repository"
)

type IThresholdService interface {
	GetSensorThreshold(sensorId int, thresholdId int) (*models.Threshold, *errors.AppError)
}

type ThresholdService struct {
	repo repository.ThresholdRepository
}

func (t ThresholdService) GetSensorThreshold(sensorId int, thresholdId int) (*models.Threshold, *errors.AppError) {
	c, err := t.repo.GetSensorThreshold(sensorId, thresholdId)
	if err != nil {
		return nil, err
	}

	//map the domain object to our dto and return it -- responsilibity of making a dto is now on the domain)

	return c, nil
}

func NewThresholdService(repo repository.ThresholdRepository) ThresholdService {
	return ThresholdService{repo}
}
