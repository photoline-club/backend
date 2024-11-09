package config

import (
	"os"

	"github.com/photoline-club/backend/database"
)

func GetDBConfig() database.DBConfig {
	return database.DBConfig{
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		DBName:   os.Getenv("MYSQL_DATABASE"),
	}
}
