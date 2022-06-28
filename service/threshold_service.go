package service

import (
	"SensorProject/middleware/errors"
	"SensorProject/models"
	"SensorProject/repository"
)

var tempCount = 5

type ThresholdService interface {
	GetSensorThreshold(sensorId uint) (*models.Threshold, *errors.AppError)
	UpsertNewThreshold(sensorId uint) (*models.Threshold, *errors.AppError)

	CheckForThresholdBreach(sensorId uint)
}

type thresholdService struct {
	ThresholdRepo      repository.ThresholdRepository
	TemperatureRepo    repository.TemperatureRepository
	ThresholdAlertRepo repository.ThresholdAlertRepository
}

func NewThresholdService(thresholdRepo repository.ThresholdRepository,
	tempRepo repository.TemperatureRepository,
	alertRepo repository.ThresholdAlertRepository) ThresholdService {
	return thresholdService{
		ThresholdRepo:      thresholdRepo,
		TemperatureRepo:    tempRepo,
		ThresholdAlertRepo: alertRepo,
	}
}

func (t thresholdService) GetSensorThreshold(sensorId uint) (*models.Threshold, *errors.AppError) {
	c, err := t.ThresholdRepo.GetSensorThreshold(sensorId)
	if err != nil {
		return nil, err
	}

	//map the domain object to our dto and return it -- responsilibity of making a dto is now on the domain)

	return c, nil
}

func (t thresholdService) UpsertNewThreshold(sensorId uint) (*models.Threshold, *errors.AppError) {
	c, err := t.ThresholdRepo.UpsertNewThreshold(sensorId)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (t thresholdService) CheckForThresholdBreach(sensorId uint) {
	threshold, err := t.ThresholdRepo.GetThresholdTemperature(sensorId)
	if err != nil {
		// TODO: log error ?
		return
	}

	temps, err := t.TemperatureRepo.GetLatestTemperatures(sensorId, tempCount)
	if err != nil {
		// TODO: log error ?
		return
	}

	for _, temp := range temps {
		if temp.LessThan(*threshold) {
			return
		}
	}

	alert := &models.ThresholdAlert{
		SensorID:    sensorId,
		Threshold:   *threshold,
		Temperature: temps[0],
	}

	t.ThresholdAlertRepo.AddThresholdAlert(alert)
}
