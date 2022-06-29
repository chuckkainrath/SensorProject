package repository

import (
	"SensorProject/middleware/errors"
	"SensorProject/models"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ThresholdRepository interface {
	GetSensorThreshold(sensorId uint, thresholdId uint) (*models.Threshold, *errors.AppError)
	PostNewThresholdToDb(thresh *models.Threshold) *errors.AppError
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

func (t thresholdRepository) GetSensorThreshold(sensorId uint, thresholdId uint) (*models.Threshold, *errors.AppError) {
	thresholdSql := "SELECT id, temperature, sensor_id FROM thresholds WHERE id = ? AND sensor_id = ?"
	var thresholds models.Threshold
	query := t.db.Raw(thresholdSql, thresholdId, sensorId)
	result := query.First(&thresholds)
	if result.Error != nil {
		return nil, errors.NewNotFoundError("Threshold Not Found")
	}
	return &thresholds, nil
}

// func (t temperatureRepository) AddTemperatureToDb(temp *models.Temperature) *errors.AppError {
// 	result := DB().Create(temp)
// 	if result.Error != nil {
// 		return errors.NewUnexpectedError("Unexpected error while processing request")
// 	}
// 	return nil
// }

func (t thresholdRepository) PostNewThresholdToDb(thresh *models.Threshold) *errors.AppError {
	//do i insert id for this if GORM autofills? TODO:DUSTIN
	result := DB().Create(&thresh)
	if result.Error != nil {
		return errors.NewUnexpectedError("Unexpected error while processing request")
	}
	return nil
	// thresholdSql := "INSERT INTO thresholds (id, temperature, sensor_id) VALUES (?,?,?)"
	// var thresholds models.Threshold
	// query := t.db.Raw(thresholdSql, sensorId)
	// result := query.Find(&thresholds)
	// if result.Error != nil {
	// 	return nil, errors.NewNotFoundError("Threshold Not Found")
	// }
	// return &thresholds, nil
}

func (t thresholdRepository) GetThresholdTemperature(sensorId uint) (*decimal.Decimal, *errors.AppError) {
	var threshold models.Threshold
	result := DB().Model(&models.Threshold{}).Select("thresholds.temperature").Where("thresholds.sensor_id = ?", sensorId).First(&threshold)
	if result.Error != nil {
		return nil, errors.NewNotFoundError("Threshold Not Found")
	}
	return &(threshold.Temperature), nil
}
