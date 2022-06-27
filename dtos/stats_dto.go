package dtos

import "time"

type InputStatsDto struct {
	SensorID uint      `mapstructure:"sensor_id"`
	From     time.Time `mapstructure:"from"`
	To       time.Time `mapstructure:"to"`
}
