package repository

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db   *gorm.DB
	once sync.Once
)

func StartPostgresDB() {
	connectionString := "postgres://sensorapp:password@localhost:5432/sensordb"
	once.Do(func() {
		var err error
		db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{

			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		fmt.Println("Connected to postgres")
	})
}

func DB() *gorm.DB { // returns a pointer to gorm.DB
	return db
}
