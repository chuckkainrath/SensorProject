package event

import (
	"SensorProject/dtos"
	"SensorProject/models"
	"SensorProject/repository"
	"SensorProject/service"
	"sync"

	"github.com/shopspring/decimal"
)

var (
	poolSize          = 10
	addTempBufferChan chan dtos.AddTemperatureDto
	once              sync.Once
	sensorMap         = make(map[uint]models.SensorThreshold)
	tempCount         = 5
)

func StartAddTemperatureHandler() {
	once.Do(func() {
		addTempBufferChan = make(chan dtos.AddTemperatureDto)

		go listenToAddTemperature()
	})
}

func GetAddTemperatureChannel() chan<- dtos.AddTemperatureDto {
	return addTempBufferChan
}

func listenToAddTemperature() {
	addTempChan := make(chan dtos.AddTemperatureDto)
	sensorTemps := make([]dtos.AddTemperatureDto, 0)

	for i := 0; i < poolSize; i++ {
		go respondToAddTemperature(addTempChan)
	}

	var output chan dtos.AddTemperatureDto
	var sensorId dtos.AddTemperatureDto
	for {
		if len(sensorTemps) > 0 {
			output = addTempChan
			sensorId = sensorTemps[0]
		}
		select {
		case sensor_id := <-addTempBufferChan:
			sensorTemps = append(sensorTemps, sensor_id)
		case output <- sensorId:
			sensorTemps = sensorTemps[1:]
			if len(sensorTemps) == 0 {
				output = nil
			}
		}
	}
}

func respondToAddTemperature(tempChan <-chan dtos.AddTemperatureDto) {
	dbClient := repository.DB()
	thresholdRepo := repository.NewThresholdRepositoryDB(dbClient)
	alertRepo := repository.NewThresholdAlertRepositoryDB(dbClient)
	tempRepo := repository.NewTemperatureRepositoryDB(dbClient)
	alertService := service.NewAlertService(alertRepo, tempRepo, thresholdRepo)

	for {
		sensorTemp := <-tempChan

		sensorData, ok := sensorMap[sensorTemp.SensorID]
		if !ok {
			sensorData, err := alertService.GetLatestTempsAndThreshold(sensorTemp.SensorID, tempCount)
			if err != nil {
				return
			}
			sensorMap[sensorTemp.SensorID] = *sensorData
		} else {
			sensorData.Temps = append(sensorData.Temps[1:], sensorTemp.Temperature)
		}

		if ok := exceedsThreshold(sensorData.Threshold, sensorData.Temps); !ok {
			alertService.AddThresholdAlert(sensorTemp.SensorID, *sensorData.Threshold, sensorTemp.Temperature)
		}
	}
}

func exceedsThreshold(threshold *decimal.Decimal, temps []decimal.Decimal) bool {
	if threshold == nil {
		return false
	}
	for _, t := range temps {
		if threshold.GreaterThan(t) {
			return false
		}
	}
	return true
}
