package config

import (
    "fmt"
    "log"
    "os"
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    "github.com/joho/godotenv"
    "golang-clinic-app/models"
)

var DB *gorm.DB

func ConnectDB() {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using environment variables")
    }

    dsn := os.Getenv("DB_URL")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    DB = db
    fmt.Println("Database connected.")

    // Auto-migrate models
    DB.AutoMigrate(&models.User{}, &models.Patient{})
}