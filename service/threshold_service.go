package service

import (
<<<<<<< HEAD
	"SensorProject/middleware/errors"
=======
>>>>>>> brooke-dev
	"SensorProject/models"
	"SensorProject/repository"
)

<<<<<<< HEAD
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
=======
var tempCount = 5

type IThresholdService interface {
	CheckForThresholdBreach(sensorId uint)
}

type thresholdService struct {
	ThresholdRepo      repository.IThresholdRepo
	TemperatureRepo    repository.ITemperatureRepo
	ThresholdAlertRepo repository.IThresholdAlertRepo
}

func NewThresholdService() IThresholdService {
	return thresholdService{
		ThresholdRepo:      repository.NewThresholdRepo(),
		TemperatureRepo:    repository.NewTemperatureRepo(),
		ThresholdAlertRepo: repository.NewThresholdAlertRepo(),
	}
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
		SensorID: sensorId,
		Threshold: *threshold,
		Temperature: temps[0],
	}

	t.ThresholdAlertRepo.AddThresholdAlert(alert)
}
>>>>>>> brooke-dev
