package repository

import (
	"SensorProject/dtos"
	"SensorProject/middleware/errors"
	"SensorProject/models"

	"gorm.io/gorm"
)

type SensorRepository interface {
	FetchSensors() (*[]dtos.SensorDto, *errors.AppError)
	FetchSensorById(sensorId uint) (*dtos.SensorDto, *errors.AppError)
}

type sensorRepository struct {
	db *gorm.DB
}

func NewSensorRepositoryDB(db *gorm.DB) SensorRepository {
	return sensorRepository{db: db}
}

// TODO: USERID
func (s sensorRepository) FetchSensors() (*[]dtos.SensorDto, *errors.AppError) {
	var sensors []dtos.SensorDto
	result := s.db.Model(&models.Sensor{}).Select("sensors.id, sensors.sensor_name as name, thresholds.temperature as threshold")
	result.Joins("Left JOIN thresholds on sensors.id = thresholds.sensor_id").Find(&sensors)
	if result.Error != nil {
		return nil, errors.NewUnexpectedError("Unexpected error while processing request")
	}
	return &sensors, nil
}

// TODO: USERID
func (s sensorRepository) FetchSensorById(sensorId uint) (*dtos.SensorDto, *errors.AppError) {
	var sensor dtos.SensorDto
	result := s.db.Model(&models.Sensor{}).Select("sensors.id, sensors.sensor_name as name, thresholds.temperature as threshold")
	result.Joins("Left JOIN thresholds on sensors.id = thresholds.sensor_id").First(&sensor, "sensors.id = ?", sensorId)
	if result.Error != nil {
		return nil, errors.NewUnexpectedError("Unexpected error while processing request")
	}
	return &sensor, nil
}
