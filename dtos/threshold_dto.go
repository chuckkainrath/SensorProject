package dtos

import "github.com/shopspring/decimal"

type AddThresholdDto struct {
	SensorID    uint            `json:"sensor_id"`
	Temperature decimal.Decimal `json:"temperature"`
}

type InputGetThresholdDto struct {
	SensorID    uint `mapstructure:"sensor_id"`
	ThresholdID uint `mapstructure:"threshold_id"`
}
