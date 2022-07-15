package service

import (
	"github.com/chuckkainrath/SensorProject/middleware/errors"
	"github.com/chuckkainrath/SensorProject/models"
	"github.com/chuckkainrath/SensorProject/repository"
	"github.com/shopspring/decimal"
)

var tempCount = 5

type ThresholdService interface {
	GetSensorThreshold(sensorId uint) (*models.Threshold, *errors.AppError)
	DeleteSensorThreshold(sensorId uint) *errors.AppError
	UpsertNewThreshold(sensorId uint, temperature decimal.Decimal) *errors.AppError

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

func (t thresholdService) UpsertNewThreshold(sensorId uint, temperature decimal.Decimal) *errors.AppError {
	//TODO:DUSTIN ???? GORM can really fill in missing ID context?
	thresh := models.Threshold{
		SensorID:    sensorId,
		Temperature: temperature,
	}
	return t.ThresholdRepo.UpsertNewThresholdToDb(&thresh)
	// c, err := t.ThresholdRepo.PostNewThreshold(sensorId)
	// if err != nil {
	// 	return nil, err
	// }
	// return c, nil
}

func (t thresholdService) DeleteSensorThreshold(sensorId uint) *errors.AppError {
	err := t.ThresholdRepo.DeleteSensorThreshold(sensorId)
	if err != nil {
		return err
	}
	return nil
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
