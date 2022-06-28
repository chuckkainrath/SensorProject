package repository

import (
	"SensorProject/dtos"
	"SensorProject/middleware/errors"

	"gorm.io/gorm"
)

type SensorsRepository interface {
	FetchSensors() ([]dtos.Sensors, *errors.AppError)
}
type sensorsRepository struct {
	db *gorm.DB
}

func NewSensorsRepositoryDB(db *gorm.DB) SensorsRepository {
	return sensorsRepository{db: db}
}

// TODO: USERID
func (s sensorsRepository) FetchSensors() ([]dtos.Sensors, *errors.AppError) {
	var sensors []dtos.Sensors
	result := s.db.Select("sensors.id, sensors.sensor_name, thresholds.temperature").Joins("Left JOIN thresholds on sensors.id = thresholds.sensor_id").Find(&sensors)
	if result.Error != nil {
		return nil, errors.NewUnexpectedError("Unexpected error while processing request")
	}
	return sensors, nil

}
