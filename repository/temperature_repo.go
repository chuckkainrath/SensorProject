package repository

import (
	"SensorProject/dtos"
	"SensorProject/models"
	"time"

	"gorm.io/gorm"
)

type ITemperatureRepo interface {
	GetPerMinuteReadingInTimeRange(sensorId uint, to, from time.Time) (*[]dtos.TemperatureDto, error)
	GetMinMaxAverageInTimeRange(sensorId uint, to, from time.Time) (*[]dtos.TemperatureStatsDto, error)
}

type temperatureRepo struct {
	db *gorm.DB
}

func NewTemperatureRepositoryDB(db *gorm.DB) ITemperatureRepo {
	return temperatureRepo{
		db: db,
	}
}

func (t temperatureRepo) GetPerMinuteReadingInTimeRange(sensorId uint, from, to time.Time) (*[]dtos.TemperatureDto, error) {
	var temps []dtos.TemperatureDto
	result := t.db.Model(&models.Temperature{}).Where("sensor_id = ? AND created_at >= ? AND created_at <= ?", sensorId, from, to)
	result.Order("created_at ASC").Find(&temps)
	if result.Error != nil {
		return nil, result.Error
	}
	return &temps, nil
}

func (t temperatureRepo) GetMinMaxAverageInTimeRange(sensorId uint, to, from time.Time) (*[]dtos.TemperatureStatsDto, error) {
	var stats []dtos.TemperatureStatsDto
	subquery := t.db.Model(&models.Temperature{}).Select("temperature as t, created_at as c")
	subquery.Where("sensor_id = ? AND created_at >= ? AND created_at <= ?", sensorId, from, to)
	result := t.db.Table("(?) as s", subquery).Select("DATE(s.c) as date, MIN(s.t) as min, MAX(s.t) as max, AVG(s.t) as average")
	result.Group("date(s.c)").Order("date(s.c) ASC").Find(&stats)
	if result.Error != nil {
		return nil, result.Error
	}
	return &stats, nil
}
