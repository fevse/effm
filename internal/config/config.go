package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	LogLevel       string
	ServHost       string
	ServPort       string
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	APIAge         string
	APISex         string
	APINationality string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		LogLevel:       os.Getenv("LOG_LEVEL"),
		ServHost:       os.Getenv("SERV_HOST"),
		ServPort:       os.Getenv("SERV_PORT"),
		DBHost:         os.Getenv("DB_HOST"),
		DBPort:         os.Getenv("DB_PORT"),
		DBUser:         os.Getenv("DB_USER"),
		DBPassword:     os.Getenv("DB_PASSWORD"),
		DBName:         os.Getenv("DB_NAME"),
		APIAge:         os.Getenv("API_AGE"),
		APISex:         os.Getenv("API_SEX"),
		APINationality: os.Getenv("API_NATIONALITY"),
	}
}

func (c *Config) DBConnectionString() string {
	return fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}
