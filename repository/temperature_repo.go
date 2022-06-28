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
	GetPerMinuteReadingInTimeRange(sensorId uint, to, from time.Time) (*[]dtos.GetTemperatureDto, *errors.AppError)
	GetMinMaxAverageInTimeRange(sensorId uint, to, from time.Time) (*[]dtos.TemperatureStatsDto, *errors.AppError)
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

func (t temperatureRepository) GetPerMinuteReadingInTimeRange(sensorId uint, from, to time.Time) (*[]dtos.GetTemperatureDto, *errors.AppError) {
	var temps []dtos.GetTemperatureDto
	result := t.db.Model(&models.Temperature{}).Where("sensor_id = ? AND created_at >= ? AND created_at <= ?", sensorId, from, to)
	result.Order("created_at ASC").Find(&temps)
	if result.Error != nil {
		return nil, errors.NewUnexpectedError("Unexpected error while processing request")
	}
	return &temps, nil
}

func (t temperatureRepository) GetMinMaxAverageInTimeRange(sensorId uint, to, from time.Time) (*[]dtos.TemperatureStatsDto, *errors.AppError) {
	var stats []dtos.TemperatureStatsDto
	subquery := t.db.Model(&models.Temperature{}).Select("temperature as t, created_at as c")
	subquery.Where("sensor_id = ? AND created_at >= ? AND created_at <= ?", sensorId, from, to)
	result := t.db.Table("(?) as s", subquery).Select("DATE(s.c) as date, MIN(s.t) as min, MAX(s.t) as max, AVG(s.t) as average")
	result.Group("date(s.c)").Order("date(s.c) ASC").Find(&stats)
	if result.Error != nil {
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
