package dtos

import (
	"time"

	"github.com/shopspring/decimal"
)

type TemperatureDto struct {
	Temperature decimal.Decimal `json:"temperature"`
	CreateAt    time.Time       `json:"created_at"`
}
