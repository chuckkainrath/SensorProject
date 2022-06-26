package repository

import (
	"SensorProject/models"

	"gorm.io/gorm"
)

type IThresholdRepo interface {
	GetSensorThreshold(sensorId string, thresholdId string) (*models.Threshold, error)
}

func (r repositoryPostgreSQL) GetSensorThreshold(sensorId string, thresholdId string) (*models.Threshold, error) {
	thresholdSql := "select id, sensor_id, temperature, from thresholds where id = ?"
	var t models.Threshold
	row := r.db.Raw(thresholdSql, thresholdId)
	result := row.Scan(&t)
	if result.Error != nil {
		return nil, result.Error
	} else {
		return nil, result.Error
	}
}

func NewThresholdRepositoryDB(dbClient *gorm.DB) IThresholdRepo {
	return repositoryPostgreSQL{dbClient}
}
