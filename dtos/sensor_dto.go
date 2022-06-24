package dtos

import "github.com/shopspring/decimal"

type Sensors struct {
	ID          uint
	Name        string
	Temperature decimal.Decimal `gorm:"type:numeric" json:"Threshold_Temperature"`
}
