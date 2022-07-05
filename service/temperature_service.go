package service

import (
	"SensorProject/dtos"
	"SensorProject/middleware/errors"
	"SensorProject/models"
	"SensorProject/repository"
	util "SensorProject/util"
	"time"

	"github.com/shopspring/decimal"
)

type TemperatureService interface {
	GetPerMinuteReading(sensorId uint, from, to time.Time, userId uint) (*[]dtos.GetTemperatureDto, *errors.AppError)
	GetMinMaxAverageStats(sensorId uint, from, to time.Time, userId uint) (*[]dtos.TemperatureStatsDto, *errors.AppError)
	AddTemperature(sensorId uint, temperature decimal.Decimal) *errors.AppError
}

type temperatureService struct {
	TemperatureRepo repository.TemperatureRepository
	DateChecker     util.DateChecker
}

func NewTemperatureService(temperatureRepo repository.TemperatureRepository, dateChecker util.DateChecker) TemperatureService {
	return temperatureService{
		TemperatureRepo: temperatureRepo,
		DateChecker:     dateChecker,
	}
}

func (t temperatureService) GetPerMinuteReading(sensorId uint, from, to time.Time, userId uint) (*[]dtos.GetTemperatureDto, *errors.AppError) {
	duration := 24 * time.Hour
	lastDay := t.DateChecker.CheckDateBeforeThresold(from, duration)
	if !lastDay {
		return nil, errors.NewBadRequestError("`from` time must be within past 24 hours")
	}
	maxDuration := t.DateChecker.CheckDateTimeDuration(from, to, duration)
	if !maxDuration {
		return nil, errors.NewBadRequestError("`from` to `to` duration must be less than 24 hours")
	}

	return t.TemperatureRepo.GetPerMinuteReadingInTimeRange(sensorId, from, to, userId)
}

func (t temperatureService) GetMinMaxAverageStats(sensorId uint, from, to time.Time, userId uint) (*[]dtos.TemperatureStatsDto, *errors.AppError) {
	duration := 30 * 24 * time.Hour
	if !t.DateChecker.CheckDateTimeDuration(from, to, duration) {
		return nil, errors.NewBadRequestError("`from` to `to` duration must be less than 30 days")
	}

	return t.TemperatureRepo.GetMinMaxAverageInTimeRange(sensorId, to, from, userId)
}

func (t temperatureService) AddTemperature(sensorId uint, temperature decimal.Decimal) *errors.AppError {
	temp := models.Temperature{
		Temperature: temperature,
		SensorID:    sensorId,
	}
	return t.TemperatureRepo.AddTemperatureToDb(&temp)
}
