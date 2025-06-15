package main

import (
    "golang-clinic-app/config"
    "golang-clinic-app/routes"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "time"
)

func main() {
    config.ConnectDB()
    router := gin.Default()

    // ✅ CORS config to allow frontend from Netlify
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"https://fantastic-pony-52cf7d.netlify.app"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    router.Static("/static", "./static")
    routes.SetupRoutes(router)

    router.Run(":8081")
}
