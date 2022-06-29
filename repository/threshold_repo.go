package repository

import (
	"SensorProject/middleware/errors"
	"SensorProject/models"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ThresholdRepository interface {
	GetSensorThreshold(sensorId uint) (*models.Threshold, *errors.AppError)
	UpsertNewThreshold(sensorId uint) (*models.Threshold, *errors.AppError)
	DeleteSensorThreshold(sensorId uint) *errors.AppError
	GetThresholdTemperature(sensorId uint) (*decimal.Decimal, *errors.AppError)
}

type thresholdRepository struct {
	db *gorm.DB
}

func NewThresholdRepositoryDB(db *gorm.DB) ThresholdRepository {
	return thresholdRepository{
		db: db,
	}
}

func (t thresholdRepository) GetSensorThreshold(sensorId uint) (*models.Threshold, *errors.AppError) {
	thresholdSql := "SELECT id, temperature, sensor_id FROM sensor_id = ?"
	var thresholds models.Threshold
	query := t.db.Raw(thresholdSql, sensorId)
	result := query.First(&thresholds)
	if result.Error != nil {
		return nil, errors.NewNotFoundError("Threshold Not Found")
	}
	return &thresholds, nil
}

func (t thresholdRepository) UpsertNewThreshold(sensorId uint) (*models.Threshold, *errors.AppError) {
	thresholdSql := "UPSERT INTO thresholds (temperature, sensor_id) VALUES (?,?)"
	var thresholds models.Threshold
	query := t.db.Raw(thresholdSql, sensorId)
	result := query.Find(&thresholds)
	if result.Error != nil {
		return nil, errors.NewNotFoundError("Threshold Not Found")
	}
	return &thresholds, nil
}

func (t thresholdRepository) DeleteSensorThreshold(sensorId uint) *errors.AppError {
	thresholdSql := "DELETE FROM thresholds (sensor_id), WHERE sensor_id = ?"
	var thresholds models.Threshold
	query := t.db.Delete(thresholdSql, sensorId)
	result := query.Where(&thresholds)
	if result.Error != nil {
		return errors.NewNotFoundError("Threshold Not Found")
	}
	return nil
}

func (t thresholdRepository) GetThresholdTemperature(sensorId uint) (*decimal.Decimal, *errors.AppError) {
	var threshold models.Threshold
	result := DB().Model(&models.Threshold{}).Select("thresholds.temperature").Where("thresholds.sensor_id = ?", sensorId).First(&threshold)
	if result.Error != nil {
		return nil, errors.NewNotFoundError("Threshold Not Found")
	}
	return &(threshold.Temperature), nil
}
