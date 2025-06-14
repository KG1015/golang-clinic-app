package main

import (
    "golang-clinic-app/config"
    "golang-clinic-app/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    config.ConnectDB()
    router := gin.Default()

    router.Static("/static", "./static")

    routes.SetupRoutes(router)
    router.Run(":8081")
}