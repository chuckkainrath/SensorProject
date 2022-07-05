package repository

import (
	"SensorProject/middleware/errors"
	"SensorProject/models"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ThresholdRepository interface {
	GetSensorThreshold(sensorId uint, userId uint) (*models.Threshold, *errors.AppError)
	UpsertNewThresholdToDb(thresh *models.Threshold) *errors.AppError
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

func (t thresholdRepository) GetSensorThreshold(sensorId uint, userId uint) (*models.Threshold, *errors.AppError) {
	thresholdSql := "SELECT thresholds.id, temperature, sensor_id FROM thresholds LEFT JOIN sensors ON thresholds.sensor_id = sensors.id WHERE sensor_id = ? AND sensors.user_id = ?"
	var thresholds models.Threshold
	query := t.db.Raw(thresholdSql, sensorId, userId)
	result := query.First(&thresholds)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, errors.NewBadRequestError("User not allowed")
	} else if result.Error != nil {
		return nil, errors.NewUnexpectedError("Unexpected error while processing request")
	}
	return &thresholds, nil
}

func (t thresholdRepository) UpsertNewThresholdToDb(thresh *models.Threshold) *errors.AppError {
	thresholdSql := "INSERT INTO thresholds (sensor_id, temperature) VALUES (?,?) on conflict(sensor_id) do update set temperature=EXCLUDED.temperature"
	var thresholds models.Threshold
	query := t.db.Raw(thresholdSql, thresh.SensorID, thresh.Temperature)
	result := query.Find(&thresholds)
	if result.Error != nil {
		return errors.NewUnexpectedError("Unexpected error while processing request")
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
