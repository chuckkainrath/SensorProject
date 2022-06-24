package dto

type Threshold struct {
	SensorID    uint    `json:"sensor_id"`
	Temperature float64 `json:"temperature"`
}
