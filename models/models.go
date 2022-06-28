package models

import (
<<<<<<< HEAD
	"time"

	"github.com/shopspring/decimal"
=======
	"github.com/shopspring/decimal"
	"time"
>>>>>>> brooke-dev
)

type User struct {
	ID             uint
	Username       string
	HashedPassword string
}

type Temperature struct {
	ID          uint
<<<<<<< HEAD
	Temperature decimal.Decimal `gorm:"type:numeric"`
=======
	Temperature decimal.Decimal
>>>>>>> brooke-dev
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
<<<<<<< HEAD
	Temperature decimal.Decimal `gorm:"type:numeric"`
=======
	Temperature decimal.Decimal
>>>>>>> brooke-dev
}

type ThresholdAlert struct {
	ID          uint
	SensorID    string
<<<<<<< HEAD
	Temperature decimal.Decimal `gorm:"type:numeric"`
	Threshold   decimal.Decimal `gorm:"type:numeric"`
=======
	Temperature decimal.Decimal
	Treshold    decimal.Decimal
>>>>>>> brooke-dev
	CreatedAt   time.Time
}
