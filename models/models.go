package models

import "time"

type User struct {
	ID             uint
	Username       string
	HashedPassword string
}

type Temperature struct {
	ID          uint
	Temperature float64
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
	Temperature float64
}

type ThresholdAlert struct {
	ID          uint
	SensorID    string
	Temperature float64
	Threshold    float64
	CreatedAt   time.Time
}
