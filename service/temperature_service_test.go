package service_test

import (
	"SensorProject/models"
	"SensorProject/service"
	"errors"
	"github.com/shopspring/decimal"
	"reflect"
	"testing"
)

// TODO: implement the service in its entirety
// TODO: figure out expetcedData

type mockTempRepo struct {
	err error
}

func (m mockTempRepo) AddTemperature(sensorId uint, temperature decimal.Decimal) error {
	return m.err
}

func TestAddTemperature(t *testing.T) {
	errRepo := errors.New("temperature repo error")
	errServer := errors.New("server error")
	inputTemperature := []models.Temperature{{ID: 1}}
	// output ?
	// outputTaskLog := []model.TaskLog{{TaskID: 1}}
	subTests := []struct {
		name         string
		temperatures []models.Temperature
		mRepo        mockTempRepo
		// expectedData *[]model.TaskLog
		expectedErr error
	}{
		{
,			name:         "SuccessTest",
			temperatures: inputTemperature,
			mRepo:        mockTempRepo{err: nil},
			//expectedData: &outputTaskLog,
			expectedErr: nil,
		},
		{
			name:         "RepoErrorTest",
			temperatures:  inputTemperature,
			mRepo:      mockTempRepo{err: errRepo},
			//expectedData: nil,
			expectedErr:  errServer,
			},
}

for _, subTest := range subTests {
	t.Run(subTest.name, func (t *testing.T) {
		temperatureService := service.TemperatureService{Temperatures: subTest.mRepo}
		result, err := temperatureService.AddTemperature(subTest.temperatures)
		if !reflect.DeepEqual(result, subTest.expectedData) {
			t.Errorf("expected (%v), got (%v)", subTest.expectedData, result)
		}
		if !errors.Is(err, subTest.expectedErr) {
			t.Errorf("expected error (%v), got error (%v)", subTest.expectedErr, err)
			}
		})
	}
}
