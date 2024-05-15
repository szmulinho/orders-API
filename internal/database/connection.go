package database

import (
	"github.com/szmulinho/orders/internal/config"
	"github.com/szmulinho/orders/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	conn := config.LoadFromEnv()
	connectionString := conn.ConnectionString()

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&model.Order{}); err != nil {
		return nil, err
	}

	return db, nil
}
