package dtos

type AddThresholdDto struct {
	SensorID    uint    `json:"sensor_id"`
	Temperature float64 `json:"temperature"`
}

type InputGetThresholdDto struct {
	SensorID    uint `mapstructure:"sensor_id"`
	ThresholdID uint `mapstructure:"threshold_id"`
}
