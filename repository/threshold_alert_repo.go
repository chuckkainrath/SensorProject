package repository

import (
	"github.com/chuckkainrath/SensorProject/middleware/errors"
	"github.com/chuckkainrath/SensorProject/models"
	"gorm.io/gorm"
)

type ThresholdAlertRepository interface {
	AddThresholdAlert(alert *models.ThresholdAlert) (err *errors.AppError)
}

type thresholdAlertRepository struct {
	db *gorm.DB
}

func NewThresholdAlertRepositoryDB(db *gorm.DB) ThresholdAlertRepository {
	return thresholdAlertRepository{
		db: db,
	}
}

func (t thresholdAlertRepository) AddThresholdAlert(alert *models.ThresholdAlert) (err *errors.AppError) {
	result := DB().Create(alert)
	if result.Error != nil {
		err = errors.NewUnexpectedError("Unexpected error while processing request")
	}
	return
}
