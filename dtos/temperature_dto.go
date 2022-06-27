package dtos

import "github.com/shopspring/decimal"

type AddTemperatureDto struct {
	SensorID    uint            `json:"sensor_id"`
	Temperature decimal.Decimal `json:"temperature" gorm:"type:numeric"`
}
