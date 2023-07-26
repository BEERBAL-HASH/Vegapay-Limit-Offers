package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// load env variables from env file
func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can't load .env file", err.Error())
	}
}

var (
	DBconn *gorm.DB
)
var PostgresHost string
var PostgresPort string
var PostgresUser string
var PostgresPass string
var PostgresName string
var PostgresSSLMode string

func init() {
	loadEnv()

	PostgresHost = os.Getenv("POSTGRES_HOST")
	PostgresPort = os.Getenv("POSTGRES_PORT")
	PostgresUser = os.Getenv("POSTGRES_USER")
	PostgresName = os.Getenv("POSTGRES_NAME")
	PostgresPass = os.Getenv("POSTGRES_PASS")
	PostgresSSLMode = os.Getenv("POSTGRES_SSLMODE")

}
