package dtos

import (
	"time"

	"github.com/shopspring/decimal"
)

// type TemperatureDto struct {
// 	Temperature decimal.Decimal `json:"temperature"`
// 	CreatedAt   time.Time       `json:"created_at"`
// }

type TemperatureStatsDto struct {
	Min     decimal.Decimal `json:"min"`
	Max     decimal.Decimal `json:"max"`
	Average decimal.Decimal `json:"average"`
	Date    time.Time       `json:"date"`
}

type AddTemperatureDto struct {
	SensorID    uint            `json:"sensor_id"`
	Temperature decimal.Decimal `json:"temperature" gorm:"type:numeric"`
}
