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

}

func NewThresholdRepositoryDB(dbClient *gorm.DB) RepositoryPostgreSQL {

	return RepositoryPostgreSQL{dbClient}
}
