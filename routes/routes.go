package routes

import (
    "github.com/gin-gonic/gin"
    "hospital-api-go-pro/controllers"
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Hospital API is running"})
    })

    r.GET("/patients", controllers.GetPatients)
    r.POST("/patients", controllers.CreatePatient)

    r.GET("/doctors", controllers.GetDoctors)
    r.POST("/doctors", controllers.CreateDoctor)
    r.GET("/doctors/schedule", controllers.GetDoctorSchedule)
}
