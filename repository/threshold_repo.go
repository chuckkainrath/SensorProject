package repository

import (
<<<<<<< HEAD
	"SensorProject/middleware/errors"
	"SensorProject/models"

	"gorm.io/gorm"
)

type ThresholdRepository interface {
	GetSensorThreshold(sensorId uint, thresholdId uint) (*models.Threshold, *errors.AppError)
	PostNewThreshold(sensorId uint) (*models.Threshold, *errors.AppError)
}

func (r repositoryPostgreSQL) GetSensorThreshold(sensorId uint, thresholdId uint) (*models.Threshold, *errors.AppError) {
	thresholdSql := "SELECT id, temperature, sensor_id FROM thresholds WHERE id = ? AND sensor_id = ?"
	var t models.Threshold
	row := r.db.Raw(thresholdSql, thresholdId, sensorId)
	result := row.Scan(&t)
	if result.Error != nil {
		return nil, errors.NewNotFoundError("Threshold Not Found")
	}
	return &t, nil
}

func (r repositoryPostgreSQL) PostNewThreshold(sensorId uint) (*models.Threshold, *errors.AppError) {
	thresholdSql := "INSERT INTO thresholds (id, temperature, sensor_id) VALUES (?,?,?)"
	var t models.Threshold
	row := r.db.Raw(thresholdSql, sensorId)
	result := row.Scan(&t)
	if result.Error != nil {
		return nil, errors.NewNotFoundError("Threshold Not Found")
	}
	return &t, nil
}

func NewThresholdRepositoryDB(dbClient *gorm.DB) ThresholdRepository {
	return repositoryPostgreSQL{dbClient}
=======
	"SensorProject/models"

	"github.com/shopspring/decimal"
)

type IThresholdRepo interface {
	GetThresholdTemperature(sensorId uint) (*decimal.Decimal, error)
}

type thresholdRepo struct{}

func NewThresholdRepo() IThresholdRepo {
	return thresholdRepo{}
}

func (t thresholdRepo) GetThresholdTemperature(sensorId uint) (*decimal.Decimal, error) {
	var threshold models.Threshold
	result := DB().Model(&models.Threshold{}).Select("thresholds.temperature").Where("thresholds.sensor_id = ?", sensorId).First(&threshold)
	if result.Error != nil {
		return nil, result.Error
	}
	return &(threshold.Temperature), nil
>>>>>>> brooke-dev
}
