package service

import (
	"SensorProject/dtos"
	"SensorProject/middleware/errors"
	"SensorProject/repository"
)

type SensorsService interface {
	GetSensorsService() ([]dtos.Sensors, *errors.AppError)
}

type sensorsService struct {
	SensorsRepository repository.SensorsRepository
}

func NewSensorsService(sensorRepo repository.SensorsRepository) SensorsService {
	return sensorsService{
		SensorsRepository: sensorRepo,
	}
}

// TODO: USER ID
func (s sensorsService) GetSensorsService() ([]dtos.Sensors, *errors.AppError) {
	return s.SensorsRepository.FetchSensors()
}
