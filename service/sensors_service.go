package service

import (
	"github.com/chuckkainrath/SensorProject/dtos"
	"github.com/chuckkainrath/SensorProject/middleware/errors"
	"github.com/chuckkainrath/SensorProject/models"
	"github.com/chuckkainrath/SensorProject/repository"
)

type SensorService interface {
	GetSensors() (*[]dtos.SensorDto, *errors.AppError)
	GetSensorById(sensorId uint) (*dtos.SensorDto, *errors.AppError)
	UpdateSensor(sensorId uint, name string, userId uint) *errors.AppError
	PostSensor(name string, userId uint) *errors.AppError
}

type sensorService struct {
	SensorsRepository repository.SensorRepository
}

func NewSensorService(sensorRepo repository.SensorRepository) SensorService {
	return sensorService{
		SensorsRepository: sensorRepo,
	}
}

// TODO: add USER ID to filter results
func (s sensorService) GetSensors() (*[]dtos.SensorDto, *errors.AppError) {
	return s.SensorsRepository.FetchSensors()
}

// TODO: add USER ID to filter results
func (s sensorService) GetSensorById(sensorId uint) (*dtos.SensorDto, *errors.AppError) {
	return s.SensorsRepository.FetchSensorById(sensorId)
}

// TODO: add USER ID to filter results
func (s sensorService) UpdateSensor(sensorId uint, name string, userId uint) *errors.AppError {

	sens := models.Sensor{
		ID:     sensorId,
		Name:   name,
		UserId: userId,
	}

	return s.SensorsRepository.UpdateSensorByID(&sens)
}

func (s sensorService) PostSensor(name string, userId uint) *errors.AppError {

	return s.SensorsRepository.CreateSensor(name, userId)

}
