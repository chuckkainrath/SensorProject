package repository

import (
	"SensorProject/dtos"
	"SensorProject/middleware/errors"
	"SensorProject/models"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type TemperatureRepository interface {
	GetPerMinuteReadingInTimeRange(sensorId uint, to, from time.Time, userId uint) (*[]dtos.GetTemperatureDto, *errors.AppError)
	GetMinMaxAverageInTimeRange(sensorId uint, to, from time.Time, userId uint) (*[]dtos.TemperatureStatsDto, *errors.AppError)
	AddTemperatureToDb(temp *models.Temperature) *errors.AppError
	GetLatestTemperatures(sensorId uint, limit int) ([]decimal.Decimal, *errors.AppError)
}

type temperatureRepository struct {
	db *gorm.DB
}

func NewTemperatureRepositoryDB(db *gorm.DB) TemperatureRepository {
	return temperatureRepository{
		db: db,
	}
}

func (t temperatureRepository) GetPerMinuteReadingInTimeRange(sensorId uint, from, to time.Time, userId uint) (*[]dtos.GetTemperatureDto, *errors.AppError) {
	var temps []dtos.GetTemperatureDto
	result := t.db.Model(&models.Temperature{}).Joins("LEFT JOIN sensors ON temperatures.sensor_id = sensors.id")
	result.Where("sensor_id = ? AND created_at >= ? AND created_at <= ? AND sensors.user_id = ?", sensorId, from, to, userId)
	result.Order("created_at ASC").Find(&temps)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, errors.NewBadRequestError("No data found")
	} else if result.Error != nil {
		return nil, errors.NewUnexpectedError("Unexpected error while processing request")
	}
	return &temps, nil
}

func (t temperatureRepository) GetMinMaxAverageInTimeRange(sensorId uint, to, from time.Time, userId uint) (*[]dtos.TemperatureStatsDto, *errors.AppError) {
	var stats []dtos.TemperatureStatsDto
	result := t.db.Model(&models.Temperature{})
	result.Select("DATE(created_at) as date, MIN(temperature) as min, MAX(temperature) as max, AVG(temperature) as average")
	result.Joins("LEFT JOIN sensors on temperatures.sensor_id = sensors.id")
	result.Where("sensor_id = ? AND created_at >= ? AND created_at <= ? AND sensors.user_id = ?", sensorId, from, to, userId)
	result.Group("date(created_at)").Order("date(created_at) ASC").Find(&stats)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, errors.NewBadRequestError("No data found")
	} else if result.Error != nil {
		return nil, errors.NewUnexpectedError("Unexpected error while processing request")
	}
	return &stats, nil
}

func (t temperatureRepository) AddTemperatureToDb(temp *models.Temperature) *errors.AppError {
	result := DB().Create(temp)
	if result.Error != nil {
		return errors.NewUnexpectedError("Unexpected error while processing request")
	}
	return nil
}

func (t temperatureRepository) GetLatestTemperatures(sensorId uint, limit int) ([]decimal.Decimal, *errors.AppError) {
	var temps []decimal.Decimal

	result := DB().Model(&models.Temperature{}).Select("temperatures.temperature").Where("temperatures.sensor_id = ?", sensorId).Order("temperatures.created_at DESC").Limit(limit).Find(&temps)

	if result.Error != nil {
		return nil, errors.NewUnexpectedError("Unexpected error while processing request")
	}
	return temps, nil
}
