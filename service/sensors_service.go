package service

import (
	"SensorProject/dtos"
	"SensorProject/repository"
)

type ISensorsService interface {
	GetSensorsService() ([]dtos.Sensors,error)
}

type sensorsService struct {
	SensorsRepository repository.ISensorsRepository
}

func NewSensorsService() ISensorsService {

	return sensorsService{
		SensorsRepository: repository.NewSensorsRepository(),
	}

}

func (s sensorsService) GetSensorsService() ([]dtos.Sensors,error) {

	fetchedSensors, err := s.SensorsRepository.FetchSensors()
	if err != nil {
		//TODO:
	}

	return fetchedSensors, nil
}
