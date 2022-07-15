package repository

import (
	"github.com/chuckkainrath/SensorProject/dtos"
	"github.com/chuckkainrath/SensorProject/middleware/errors"
	"github.com/chuckkainrath/SensorProject/models"
	"gorm.io/gorm"
)

type SensorRepository interface {
	FetchSensors() (*[]dtos.SensorDto, *errors.AppError)
	FetchSensorById(sensorId uint) (*dtos.SensorDto, *errors.AppError)
	UpdateSensorByID(sens *models.Sensor) *errors.AppError
	CreateSensor(name string, userId uint) *errors.AppError
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

func (s sensorRepository) UpdateSensorByID(sens *models.Sensor) *errors.AppError {

	sensorSql := "UPDATE sensors SET sensor_name = ? WHERE id = ?"

	var sensor models.Sensor

	query := s.db.Raw(sensorSql, sens.Name, sens.ID)
	result := query.Find(&sensor)

	if result.Error != nil {
		return errors.NewUnexpectedError("Unexpected error while processing request")
	}
	return nil
}

func (s sensorRepository) CreateSensor(name string, userId uint) *errors.AppError {

	sensorSql := "INSERT INTO sensors (user_id, sensor_name) VALUES (?, ?)"

	// var sensor models.Sensor

	result := s.db.Exec(sensorSql, userId, name)

	if result.Error != nil {
		return errors.NewUnexpectedError("Unexpected error while processing request")
	}
	return nil

}
