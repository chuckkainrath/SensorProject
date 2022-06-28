package dtos

import "github.com/shopspring/decimal"

type SensorDto struct {
	ID          uint            `json:"id,omitempty"`
	Name        string          `json:"name"`
	Temperature decimal.Decimal `gorm:"type:numeric" json:"temperature"`
}

type GetSensorDto struct {
	SensorID uint `json:"sensor_id"`
}

type UpdateSensorDto struct {
	SensorID uint   `json:"sensor_id"`
	Name     string `json:"name"`
}
