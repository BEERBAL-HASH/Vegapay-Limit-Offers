package database

import (
	"fmt"
	"log"

	"github.com/BEERBAL-HASH/Vegapay-Limit-Offer/config"
	"github.com/BEERBAL-HASH/Vegapay-Limit-Offer/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=%s", config.PostgresHost, config.PostgresUser, config.PostgresPass, config.PostgresName, config.PostgresSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database server! \n", err.Error())
		panic(err)
	}
	fmt.Println("Database connected successfully")
	db.AutoMigrate(&models.Account{}, &models.LimitOffer{})
	config.DBconn = db
}
