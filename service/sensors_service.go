package service

import (
	"SensorProject/dtos"
	"SensorProject/middleware/errors"
	"SensorProject/repository"
)

type SensorService interface {
	GetSensors() (*[]dtos.SensorDto, *errors.AppError)
	GetSensorById(sensorId uint) (*dtos.SensorDto, *errors.AppError)
	UpdateSensor()
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
func (s sensorService) UpdateSensor() {

}
