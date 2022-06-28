package service

import (
	"SensorProject/dtos"
<<<<<<< HEAD
	"SensorProject/middleware/errors"
	"SensorProject/repository"
	util "SensorProject/util"
	"time"
)

type ITemperatureService interface {
	GetPerMinuteReading(sensorId uint, from, to time.Time) (*[]dtos.TemperatureDto, *errors.AppError)
	GetMinMaxAverageStats(sensorId uint, from, to time.Time) (*[]dtos.TemperatureStatsDto, *errors.AppError)
=======
	"SensorProject/models"
	"SensorProject/repository"
)

type ITemperatureService interface {
	AddTemperature(tempDto dtos.AddTemperatureDto) error
>>>>>>> brooke-dev
}

type temperatureService struct {
	TemperatureRepo repository.ITemperatureRepo
<<<<<<< HEAD
	DateChecker     util.IDateChecker
}

func NewTemperatureService(temperatureRepo repository.ITemperatureRepo, dateChecker util.IDateChecker) ITemperatureService {
	return temperatureService{
		TemperatureRepo: temperatureRepo,
		DateChecker:     dateChecker,
	}
}

func (t temperatureService) GetPerMinuteReading(sensorId uint, from, to time.Time) (*[]dtos.TemperatureDto, *errors.AppError) {
	duration := 24 * time.Hour
	lastDay := t.DateChecker.CheckDateBeforeThresold(from, duration)
	if !lastDay {
		return nil, errors.NewBadRequestError("`from` time must be within past 24 hours")
	}
	maxDuration := t.DateChecker.CheckDateTimeDuration(from, to, duration)
	if !maxDuration {
		return nil, errors.NewBadRequestError("`from` to `to` duration must be less than 24 hours")
	}

	return t.TemperatureRepo.GetPerMinuteReadingInTimeRange(sensorId, from, to)
}

func (t temperatureService) GetMinMaxAverageStats(sensorId uint, from, to time.Time) (*[]dtos.TemperatureStatsDto, *errors.AppError) {
	duration := 30 * 24 * time.Hour
	if !t.DateChecker.CheckDateTimeDuration(from, to, duration) {
		return nil, errors.NewBadRequestError("`from` to `to` duration must be less than 30 days")
	}

	return t.TemperatureRepo.GetMinMaxAverageInTimeRange(sensorId, to, from)
=======
}

func NewTemperatureService() ITemperatureService {
	return temperatureService{
		TemperatureRepo: repository.NewTemperatureRepo(),
	}
}

func (t temperatureService) AddTemperature(tempDto dtos.AddTemperatureDto) error {

	temp := models.Temperature{
		Temperature: tempDto.Temperature,
		SensorID:    tempDto.SensorID,
	}
	return t.TemperatureRepo.AddTemperatureToDb(&temp)

>>>>>>> brooke-dev
}
