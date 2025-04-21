package db

import (
	"fmt"
	"log"

	"khajuraho/backend/config"
	"khajuraho/backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Instance *gorm.DB

func Connect() {
	user := config.AppConfig.DBUser
	password := config.AppConfig.DBPassword
	name := config.AppConfig.DBName
	host := config.AppConfig.DBHost
	port := config.AppConfig.DBPort

	if user == "" || password == "" || name == "" || host == "" {
		log.Fatal("database credentials are not set properly in environment variables")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		host, user, password, name, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	log.Println("✅ Successfully connected to the database")

	Instance = db

	migrate()
}

// TODO: add disconnect func.
func Disconnect() {
	db, err := Instance.DB()
	if err != nil {
		log.Printf("❌ Failed to get raw database connection: %v", err)
		return
	}

	if err := db.Close(); err != nil {
		log.Printf("❌ Error while closing database: %v", err)
		return
	}

	log.Println("🔌 Database connection closed gracefully")
}

func migrate() {
	err := Instance.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		log.Fatalf("❌ AutoMigrate failed: %v", err)
	}

	log.Println("📦 Database tables migrated (created if not exists)")
}
