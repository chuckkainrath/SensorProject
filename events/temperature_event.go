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
	poolSize                  = 10
	addTempBufferChan         chan dtos.AddTemperatureDto
	updateThresholdBufferChan chan dtos.ThresholdEventDto
	once                      sync.Once
	sensorMap                 = make(map[uint]models.SensorThreshold)
	tempCount                 = 5
)

func StartAddTemperatureHandler() {
	once.Do(func() {
		addTempBufferChan = make(chan dtos.AddTemperatureDto)
		updateThresholdBufferChan = make(chan dtos.ThresholdEventDto)
		go listenToAddTemperature()
		go listenToThresholdUpdate()
	})
}

func GetAddTemperatureChannel() chan<- dtos.AddTemperatureDto {
	return addTempBufferChan
}

func GetUpdateThresholdChannel() chan<- dtos.ThresholdEventDto {
	return updateThresholdBufferChan
}

func listenToAddTemperature() {
	addTempChan := make(chan dtos.AddTemperatureDto)
	sensorTemps := make([]dtos.AddTemperatureDto, 0)

	for i := 0; i < poolSize; i++ {
		go respondToAddTemperature(addTempChan)
	}

	var tempOutput chan dtos.AddTemperatureDto
	var tempSensorDto dtos.AddTemperatureDto

	for {
		if len(sensorTemps) > 0 {
			tempOutput = addTempChan
			tempSensorDto = sensorTemps[0]
		}
		select {
		case tempDto := <-addTempBufferChan:
			sensorTemps = append(sensorTemps, tempDto)
		case tempOutput <- tempSensorDto:
			sensorTemps = sensorTemps[1:]
			if len(sensorTemps) == 0 {
				tempOutput = nil
			}
		}
	}
}

func listenToThresholdUpdate() {
	updateThresholdChan := make(chan dtos.ThresholdEventDto)
	thresholds := make([]dtos.ThresholdEventDto, 0)

	for i := 0; i < poolSize; i++ {
		go respondToUpdateThreshold(updateThresholdChan)
	}

	var thresholdOutput chan dtos.ThresholdEventDto
	var thresholdSensorDto dtos.ThresholdEventDto
	for {
		if len(thresholds) > 0 {
			thresholdOutput = updateThresholdChan
			thresholdSensorDto = thresholds[0]
		}
		select {
		case thresholdDto := <-updateThresholdBufferChan:
			thresholds = append(thresholds, thresholdDto)
		case thresholdOutput <- thresholdSensorDto:
			thresholds = thresholds[1:]
			if len(thresholds) == 0 {
				thresholdOutput = nil
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
			sensorDataRes, err := alertService.GetLatestTempsAndThreshold(sensorTemp.SensorID, tempCount)
			if err != nil {
				return
			}
			sensorMap[sensorTemp.SensorID] = *sensorDataRes
			sensorData = *sensorDataRes
		} else {
			sensorData.Temps = append([]decimal.Decimal{sensorTemp.Temperature}, sensorData.Temps[0:(tempCount-1)]...)
		}
		if exceeded := exceedsThreshold(sensorData.Threshold, sensorData.Temps); exceeded {
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

func respondToUpdateThreshold(updateThresholdChan <-chan dtos.ThresholdEventDto) {
	for {
		newThreshold := <-updateThresholdChan

		sensorData := sensorMap[newThreshold.SensorID]
		sensorData.Threshold = newThreshold.Temperature
	}
}
