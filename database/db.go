package db

import (
	"fmt"
	"log"

	"khajuraho/backend/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	Db *gorm.DB
}

var DB DBInstance

func ConnectDB() error {
	user := config.AppConfig.DBUser
	password := config.AppConfig.DBPassword
	name := config.AppConfig.DBName
	host := config.AppConfig.DBHost
	port := config.AppConfig.DBPort

	if user == "" || password == "" || name == "" || host == "" {
		return fmt.Errorf("database credentials are not set properly in environment variables")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		host, user, password, name, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	log.Println("✅ Successfully connected to the database")

	DB = DBInstance{Db: db}
	return nil
}
