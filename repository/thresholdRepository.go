package repository

import (
	"SensorProject/middleware/errors"
	"SensorProject/models"

	"gorm.io/gorm"
)

type ThresholdRepository interface {
	GetSensorThreshold(sensorId string, thresholdId string) (*models.Threshold, *errors.AppError)
}

func (r RepositoryPostgreSQL) GetSensorThreshold(sensorId string, thresholdId string) (*models.Threshold, *errors.AppError) {
	thresholdSql := "select id, sensor_id, temperature, from thresholds where id = ?"
	var t models.Threshold
	row := r.db.Raw(thresholdSql, thresholdId)
	err := row.Scan(&t)
	if err != nil {
		return nil, errors.NewNotFoundError("Threshold Not Found")
	} else {
		return nil, errors.NewUnexpectedError("unexpected database error")
	}

}

func NewThresholdRepositoryDB(dbClient *gorm.DB) RepositoryPostgreSQL {

	return RepositoryPostgreSQL{dbClient}
}
