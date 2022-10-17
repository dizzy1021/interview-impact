package driver

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"dizzy1021.dev/interview-impact/util"
)

func LoadDatabase() *gorm.DB {
	
	dbHost := util.LoadEnv("DB_HOST")
	dbPort := util.LoadEnv("DB_PORT")
	dbUser := util.LoadEnv("DB_USER")
	dbPass := util.LoadEnv("DB_PASSWORD")
	dbName := util.LoadEnv("DB_DATABASE")
	dbSSLMode := util.LoadEnv("DB_SSLMODE")
	dbTimezone := util.LoadEnv("DB_TIMEZONE")
	
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", 
		dbHost,
		dbUser,
		dbPass,
		dbName,
		dbPort,						
		dbSSLMode,
		dbTimezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error to load database")
	}
	return db
}