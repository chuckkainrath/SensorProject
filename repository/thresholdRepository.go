package repository

import (
	"SensorProject/middleware/errors"
	"SensorProject/models"

	"gorm.io/gorm"
)

type ThresholdRepository interface {
	GetSensorThreshold(sensorId int, thresholdId int) (*models.Threshold, *errors.AppError)
	PostNewThreshold(sensorId int) (*models.Threshold, *errors.AppError)
}

func (r RepositoryPostgreSQL) GetSensorThreshold(sensorId int, thresholdId int) (*models.Threshold, *errors.AppError) {
	thresholdSql := "SELECT id, temperature, sensor_id FROM thresholds WHERE id = ? AND sensor_id = ?"
	var t models.Threshold
	row := r.db.Raw(thresholdSql, thresholdId, sensorId)
	result := row.Scan(&t)
	if result.Error != nil {
		return nil, errors.NewNotFoundError("Threshold Not Found")
	}
	return &t, nil
	// } else {
	// 	return nil, errors.NewUnexpectedError("unexpected database error")
	// }

}

func (r RepositoryPostgreSQL) PostNewThreshold(sensorId int) (*models.Threshold, *errors.AppError) {
	thresholdSql := "INSERT INTO thresholds (id, temperature, sensor_id) VALUES (?,?,?)"
	var t models.Threshold
	row := r.db.Raw(thresholdSql, sensorId)
	result := row.Scan(&t)
	if result.Error != nil {
		return nil, errors.NewNotFoundError("Threshold Not Found")
	}
	return &t, nil
}

func NewThresholdRepositoryDB(dbClient *gorm.DB) RepositoryPostgreSQL {

	return RepositoryPostgreSQL{dbClient}
}
