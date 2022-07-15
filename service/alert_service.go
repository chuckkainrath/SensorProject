package service

import (
	"github.com/chuckkainrath/SensorProject/middleware/errors"
	"github.com/chuckkainrath/SensorProject/models"
	"github.com/chuckkainrath/SensorProject/repository"
	"github.com/shopspring/decimal"
)

type AlertService interface {
	AddThresholdAlert(sensorId uint, threshold, temperature decimal.Decimal)
	GetLatestTempsAndThreshold(sensorId uint, tempCount int) (*models.SensorThreshold, *errors.AppError)
	GetThresholdTemperature(sensorId uint) (*decimal.Decimal, *errors.AppError)
}

type alertService struct {
	ThresholdAlertRepo repository.ThresholdAlertRepository
	TemperatureRepo    repository.TemperatureRepository
	ThresholdRepo      repository.ThresholdRepository
}

func NewAlertService(alertRepo repository.ThresholdAlertRepository, tempRepo repository.TemperatureRepository, thresholdRepo repository.ThresholdRepository) AlertService {
	return alertService{
		ThresholdAlertRepo: alertRepo,
		TemperatureRepo:    tempRepo,
		ThresholdRepo:      thresholdRepo,
	}
}

func (a alertService) AddThresholdAlert(sensorId uint, threshold, temperature decimal.Decimal) {
	alert := &models.ThresholdAlert{
		SensorID:    sensorId,
		Threshold:   threshold,
		Temperature: temperature,
	}

	a.ThresholdAlertRepo.AddThresholdAlert(alert)
}

func (a alertService) GetLatestTempsAndThreshold(sensorId uint, tempCount int) (*models.SensorThreshold, *errors.AppError) {
	sensorThreshold := models.SensorThreshold{}
	temps, err := a.TemperatureRepo.GetLatestTemperatures(sensorId, tempCount)
	if err != nil {
		return nil, err
	}
	sensorThreshold.Temps = temps

	threshold, _ := a.ThresholdRepo.GetThresholdTemperature(sensorId)
	sensorThreshold.Threshold = threshold

	return &sensorThreshold, nil
}

func (a alertService) GetThresholdTemperature(sensorId uint) (*decimal.Decimal, *errors.AppError) {
	return a.ThresholdRepo.GetThresholdTemperature(sensorId)
}
