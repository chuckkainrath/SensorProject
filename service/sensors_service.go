package service

import (
	"SensorProject/dtos"
	"SensorProject/repository"
)

type ISensorsService interface {
	GetSensorsService() []dtos.Sensors
	GetSensorIdService(sensorId int) (dtos.Sensors, error)
	UpdateSensor()
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

func (s sensorsService) GetSensorIdService(sensorId int) (dtos.Sensors, error) {
	fetchedSensorIds, err := s.SensorsRepository.FetchSensorId(sensorId)
	if err != nil {
		//TODO:
	}
	return fetchedSensorIds, nil
}

func (s sensorsService) UpdateSensor() {

}
