package service

import (
	"SensorProject/dtos"
	"SensorProject/repository"
	util "SensorProject/util"
	"errors"
	"time"
)

type ITemperatureService interface {
	GetPerMinuteReading(sensorId uint, from, to time.Time) (*[]dtos.TemperatureDto, error)
	GetMinMaxAverageStats(sensorId uint, from, to time.Time) (*[]dtos.TemperatureStatsDto, error)
}

type temperatureService struct {
	TemperatureRepo repository.ITemperatureRepo
	DateChecker     util.IDateChecker
}

func NewTemperatureService(temperatureRepo repository.ITemperatureRepo, dateChecker util.IDateChecker) ITemperatureService {
	return temperatureService{
		TemperatureRepo: temperatureRepo,
		DateChecker:     dateChecker,
	}
}

func (t temperatureService) GetPerMinuteReading(sensorId uint, from, to time.Time) (*[]dtos.TemperatureDto, error) {
	duration := 24 * time.Hour
	lastDay := t.DateChecker.CheckDateBeforeThresold(from, duration)
	maxDuration := t.DateChecker.CheckDateTimeDuration(from, to, duration)
	if !lastDay || !maxDuration {
		return nil, errors.New("time error") // TODO: Do a more formal error checking/response
	}

	return t.TemperatureRepo.GetPerMinuteReadingInTimeRange(sensorId, from, to)
}

func (t temperatureService) GetMinMaxAverageStats(sensorId uint, from, to time.Time) (*[]dtos.TemperatureStatsDto, error) {
	duration := 30 * 24 * time.Hour
	if !t.DateChecker.CheckDateTimeDuration(from, to, duration) {
		return nil, errors.New("time error") // TODO: Do a more formal error checking/response
	}

	return t.TemperatureRepo.GetMinMaxAverageInTimeRange(sensorId, to, from)
}
