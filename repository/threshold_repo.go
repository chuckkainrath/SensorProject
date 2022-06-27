package repository

import (
	"SensorProject/middleware/errors"
	"SensorProject/models"

	"gorm.io/gorm"
)

type IThresholdRepo interface {
	GetSensorThreshold(sensorId uint, thresholdId uint) (*models.Threshold, *errors.AppError)
}

func (r repositoryPostgreSQL) GetSensorThreshold(sensorId uint, thresholdId uint) (*models.Threshold, *errors.AppError) {
	thresholdSql := "select id, sensor_id, temperature, from thresholds where id = ?"
	var t models.Threshold
	row := r.db.Raw(thresholdSql, thresholdId)
	result := row.Scan(&t)
	if result.Error != nil {
		return nil, errors.NewNotFoundError("Threshold Not Found")
	} else {
		return nil, errors.NewUnexpectedError("unexpected database error")
	}
}

func NewThresholdRepositoryDB(dbClient *gorm.DB) IThresholdRepo {
	return repositoryPostgreSQL{dbClient}
}
