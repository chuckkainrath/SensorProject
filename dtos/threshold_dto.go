package dtos

type AddThresholdDto struct {
	SensorID    uint    `json:"sensor_id"`
	Temperature float64 `json:"temperature"`
}

type InputGetThresholdDto struct {
	SensorID    uint `json:"sensor_id"`
	ThresholdID uint `json:"threshold_id"`
}
