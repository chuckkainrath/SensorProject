package service

import (
	"SensorProject/models"
	"SensorProject/repository"
)

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
