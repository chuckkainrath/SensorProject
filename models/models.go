package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type User struct {
	ID             uint
	Username       string
	HashedPassword string
}

type Temperature struct {
	ID          uint
	Temperature decimal.Decimal `gorm:"type:numeric"`
	SensorID    uint
	CreatedAt   time.Time
}

type Sensor struct {
	ID     uint
	Name   string
	UserId uint
}

type Threshold struct {
	ID          uint
	SensorID    uint
	Temperature decimal.Decimal `gorm:"type:numeric"`
}

type ThresholdAlert struct {
	ID          uint
	SensorID    string
	Temperature decimal.Decimal `gorm:"type:numeric"`
	Threshold   decimal.Decimal `gorm:"type:numeric"`
	CreatedAt   time.Time
}
