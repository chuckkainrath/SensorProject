package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type User struct {
	ID             uint
	Username       string
	HashedPassword string
}

type Temperature struct {
	ID          uint
	Temperature decimal.Decimal
	SensorID    uint
	CreatedAt   time.Time
}

type Sensor struct {
	ID     uint
	Name   string
	UserId uint
}

type Treshold struct {
	ID          uint
	SensorID    uint
	Temperature decimal.Decimal
}

type TresholdAlert struct {
	ID          uint
	SensorID    string
	Temperature decimal.Decimal
	Treshold    decimal.Decimal
	CreatedAt   time.Time
}
