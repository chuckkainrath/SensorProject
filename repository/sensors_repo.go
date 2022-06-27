package repository

import (
	"SensorProject/dtos"
)

type ISensorsRepository interface {
	FetchSensors() []dtos.Sensors
	FetchSensorId(sensorId int) (dtos.Sensors, error)
}
type sensorsRepository struct{}

func NewSensorsRepository() ISensorsRepository {
	return sensorsRepository{}
}

func (s sensorsRepository) FetchSensors() []dtos.Sensors {
	var sensors []dtos.Sensors
	result := DB().Select("sensors.id,sensors.sensor_name,thresholds.temperature").Joins("Left JOIN thresholds on sensors.id = thresholds.sensor_id").Find(&sensors)
	if result.Error != nil {

	}
	return sensors
}

func (s sensorsRepository) FetchSensorId(sensorId int) (dtos.Sensors, error) {

	var sensors dtos.Sensors
	result := DB().Select("sensors.id,sensors.sensor_name,thresholds.temperature").Joins("JOIN thresholds on sensors.id = thresholds.sensor_id").Find(&sensors, "sensor_id = ?", sensorId)
	if result.Error != nil {

	}
	return sensors, nil
}
