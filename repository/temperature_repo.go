package repository

import (
<<<<<<< HEAD
	"SensorProject/dtos"
	"SensorProject/middleware/errors"
	"SensorProject/models"
	"time"

	"gorm.io/gorm"
)

type ITemperatureRepo interface {
	GetPerMinuteReadingInTimeRange(sensorId uint, to, from time.Time) (*[]dtos.TemperatureDto, *errors.AppError)
	GetMinMaxAverageInTimeRange(sensorId uint, to, from time.Time) (*[]dtos.TemperatureStatsDto, *errors.AppError)
}

type temperatureRepo struct {
	db *gorm.DB
}

func NewTemperatureRepositoryDB(db *gorm.DB) ITemperatureRepo {
	return temperatureRepo{
		db: db,
	}
}

func (t temperatureRepo) GetPerMinuteReadingInTimeRange(sensorId uint, from, to time.Time) (*[]dtos.TemperatureDto, *errors.AppError) {
	var temps []dtos.TemperatureDto
	result := t.db.Model(&models.Temperature{}).Where("sensor_id = ? AND created_at >= ? AND created_at <= ?", sensorId, from, to)
	result.Order("created_at ASC").Find(&temps)
	if result.Error != nil {
		return nil, errors.NewUnexpectedError("Unexpected error while processing request")
	}
	return &temps, nil
}

func (t temperatureRepo) GetMinMaxAverageInTimeRange(sensorId uint, to, from time.Time) (*[]dtos.TemperatureStatsDto, *errors.AppError) {
	var stats []dtos.TemperatureStatsDto
	subquery := t.db.Model(&models.Temperature{}).Select("temperature as t, created_at as c")
	subquery.Where("sensor_id = ? AND created_at >= ? AND created_at <= ?", sensorId, from, to)
	result := t.db.Table("(?) as s", subquery).Select("DATE(s.c) as date, MIN(s.t) as min, MAX(s.t) as max, AVG(s.t) as average")
	result.Group("date(s.c)").Order("date(s.c) ASC").Find(&stats)
	if result.Error != nil {
		return nil, errors.NewUnexpectedError("Unexpected error while processing request")
	}
	return &stats, nil
=======
	"SensorProject/models"

	"github.com/shopspring/decimal"
)

type ITemperatureRepo interface {
	AddTemperatureToDb(temp *models.Temperature) error
	GetLatestTemperatures(sensorId uint, limit int) ([]decimal.Decimal, error)
}

type temperatureRepo struct{}

func NewTemperatureRepo() ITemperatureRepo {
	return temperatureRepo{}
}

func (t temperatureRepo) AddTemperatureToDb(temp *models.Temperature) error {
	result := DB().Create(temp)
	return result.Error
}

func (t temperatureRepo) GetLatestTemperatures(sensorId uint, limit int) ([]decimal.Decimal, error) {
	var temps []decimal.Decimal

	result := DB().Model(&models.Temperature{}).Select("temperatures.temperature").Where("temperatures.sensor_id = ?", sensorId).Order("temperatures.created_at DESC").Limit(limit).Find(&temps)

	if result.Error != nil {
		return nil, result.Error
	}
	return temps, nil
>>>>>>> brooke-dev
}
