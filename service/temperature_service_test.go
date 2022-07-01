package service

import (
	"SensorProject/dtos"
	"SensorProject/middleware/errors"
	"SensorProject/models"
	"SensorProject/util"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTemperatureRepo struct {
	mock.Mock
}

func (mock *MockTemperatureRepo) GetPerMinuteReadingInTimeRange(sensorId uint, to, from time.Time) (*[]dtos.GetTemperatureDto, *errors.AppError) {
	return nil, nil
}

func (mock *MockTemperatureRepo) GetMinMaxAverageInTimeRange(sensorId uint, to, from time.Time) (*[]dtos.TemperatureStatsDto, *errors.AppError) {
	return nil, nil
}

func (mock *MockTemperatureRepo) AddTemperatureToDb(temp *models.Temperature) *errors.AppError {

	if temp.SensorID == 0 || temp.Temperature == decimal.NewFromFloat32(0.0) {
		return errors.NewUnexpectedError("Unexpected error while processing request")

	}
	return nil

}

func (mock *MockTemperatureRepo) GetLatestTemperatures(sensorId uint, limit int) ([]decimal.Decimal, *errors.AppError) {

	return nil, nil
}

func TestAddTemperatureNoError(t *testing.T) {
	mockRepo := new(MockTemperatureRepo)

	temp := models.NewTemperatureModel()
	NewTempModel = func() models.Temperature {
		return temp
	}

	id := uint(1)
	Temperature := decimal.NewFromFloat32(70.2)

	temp.SensorID = id
	temp.Temperature = Temperature

	testTemperatureService := NewTemperatureService(mockRepo, util.NewDateChecker())

	mockRepo.On("AddTemperatureToDb", &temp).Return(nil)

	if NewTempModel().SensorID != id || NewTempModel().Temperature != Temperature {
		t.Errorf("Expected SensorID (%v), got (%v)\n", id, temp.SensorID)
	}

	err := testTemperatureService.AddTemperature(id, Temperature)

	assert.Nil(t, err)



}


func TestAddTemperatureError(t *testing.T) {
	mockRepo := new(MockTemperatureRepo)

	temp := models.NewTemperatureModel()
	NewTempModel = func() models.Temperature {
		return temp
	}

	id := uint(0)
	Temperature := decimal.NewFromFloat32(70.2)

	temp.SensorID = id
	temp.Temperature = Temperature

	testTemperatureService := NewTemperatureService(mockRepo, util.NewDateChecker())

	mockRepo.On("AddTemperatureToDb", &temp).Return(nil)

	if NewTempModel().SensorID != id || NewTempModel().Temperature != Temperature {
		t.Errorf("Expected SensorID (%v), got (%v)\n", id, temp.SensorID)
	}

	err := testTemperatureService.AddTemperature(id, Temperature)


	assert.Equal(t, errors.NewUnexpectedError("Unexpected error while processing request"), err)

}
