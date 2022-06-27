package dtos

type InputStatsDto struct {
	SensorID uint      `json:"sensor_id"`
	From     time.Time `json:"from"`
	To       time.Time `json:"to"`
}
