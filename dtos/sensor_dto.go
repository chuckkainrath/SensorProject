package dtos

import "github.com/shopspring/decimal"

type SensorDto struct {
	ID        uint            `json:"id,omitempty"`
	Name      string          `json:"name"`
	Threshold decimal.Decimal `gorm:"type:numeric" json:"threshold"`
}

type SensorIdDto struct {
	SensorID uint `mapstructure:"sensor_id"`
}

type UpdateSensorDto struct {
	SensorID uint   `json:"sensor_id"`
	Name     string `json:"name"`
	UserID   uint   `json:"user_id"`
}

type PostSensorDto struct {
	Name string `json:"name"`
}
