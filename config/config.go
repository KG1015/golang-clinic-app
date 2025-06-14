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
    // Load .env in development only (Render ignores this)
    _ = godotenv.Load()

    dsn := os.Getenv("DATABASE_URL") // ✅ use correct env var
    if dsn == "" {
        log.Fatal("DATABASE_URL is not set")
    }

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    DB = db
    fmt.Println("Database connected.")

    // Auto-migrate your models
    DB.AutoMigrate(&models.User{}, &models.Patient{})
}
