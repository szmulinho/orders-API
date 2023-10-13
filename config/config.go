package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type StorageConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	Port     string `json:"port"`
}

func LoadConfigFromEnv() StorageConfig {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	return StorageConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}
}

func (c StorageConfig) ConnectionString() string {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		c.Host, c.User, c.Password, c.Dbname, c.Port)
	return connectionString
}
