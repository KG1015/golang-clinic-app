package routes

import (
    "github.com/gin-gonic/gin"
    "golang-clinic-app/controllers"
    "golang-clinic-app/middleware"
)

func SetupRoutes(router *gin.Engine) {
    api := router.Group("/api")
    {
        api.POST("/login", controllers.Login)
        api.POST("/register", controllers.RegisterUser)

        // Protected routes (require JWT)
        protected := api.Group("/")
        protected.Use(middleware.JWTAuth())
        {
			protected.POST("/patients", controllers.CreatePatient)
            protected.GET("/patients", controllers.GetPatients)
            protected.GET("/patients/:id", controllers.GetPatient)
            protected.PUT("/patients/:id", controllers.UpdatePatient)
            protected.DELETE("/patients/:id", controllers.DeletePatient)
        }
    }
}