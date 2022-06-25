package service

import (
	"SensorProject/middleware/errors"
	"SensorProject/models"
	"SensorProject/repository"
)

type IThresholdService interface {
	GetSensorThreshold(sensorId string, thresholdId string) (*models.Threshold, *errors.AppError)
}

type thresholdService struct {
	thresholdRepo repository.IThresholdRepo
}

func NewThresholdService(repo repository.IThresholdRepo) IThresholdService {
	return thresholdService{thresholdRepo:repo}
}

func (t thresholdService) GetSensorThreshold(sensorId string, thresholdId string) (*models.Threshold, *errors.AppError) {
	c, err := t.thresholdRepo.GetSensorThreshold(sensorId, thresholdId)
	if err != nil {
		return nil, err
	}

	//map the domain object to our dto and return it -- responsilibity of making a dto is now on the domain)

	return c, nil
}


