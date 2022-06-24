package service

import (
	"SensorProject/dtos"
	"SensorProject/repository"
)

type ISensorsService interface {
	GetSensorsService() []dtos.Sensors
}

type sensorsService struct {
	SensorsRepository repository.ISensorsRepository
}

func NewSensorsService() ISensorsService {

	return sensorsService{
		SensorsRepository: repository.NewSensorsRepository(),
	}

}

func (s sensorsService) GetSensorsService() []dtos.Sensors {

	fetchedSensors := s.SensorsRepository.FetchSensors()

	return fetchedSensors
}
