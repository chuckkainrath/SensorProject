package dtos

type AddTemperatureDto struct {
	SensorID    uint    `json:"sensor_id"`
	Temperature float32 `json:"temperature"`
}
