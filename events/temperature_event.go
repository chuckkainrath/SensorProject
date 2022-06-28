package event

import (
	"SensorProject/repository"
	"SensorProject/service"
	"sync"
)

var (
	poolSize          = 10
	addTempBufferChan chan uint
	once              sync.Once
)

func StartAddTemperatureHandler() {
	once.Do(func() {
		addTempBufferChan = make(chan uint)

		go listenToAddTemperature()
	})
}

func GetAddTemperatureChannel() chan<- uint {
	return addTempBufferChan
}

func listenToAddTemperature() {
	addTempChan := make(chan uint)
	sensors := make([]uint, 0)

	for i := 0; i < poolSize; i++ {
		go respondToAddTemperature(addTempChan)
	}

	var output chan uint
	var sensorId uint
	for {
		if len(sensors) > 0 {
			output = addTempChan
			sensorId = sensors[0]
		}
		select {
		case sensor_id := <-addTempBufferChan:
			sensors = append(sensors, sensor_id)
		case output <- sensorId:
			sensors = sensors[1:]
			if len(sensors) == 0 {
				output = nil
			}
		}
	}
}

func respondToAddTemperature(tempChan <-chan uint) {
	dbClient := repository.DB()
	thresholdRepo := repository.NewThresholdRepositoryDB(dbClient)
	tempRepo := repository.NewTemperatureRepositoryDB(dbClient)
	alertRepo := repository.NewThresholdAlertRepositoryDB(dbClient)
	thresholdService := service.NewThresholdService(thresholdRepo, tempRepo, alertRepo)

	for {
		sensorId := <-tempChan
		thresholdService.CheckForThresholdBreach(sensorId)
	}
}
