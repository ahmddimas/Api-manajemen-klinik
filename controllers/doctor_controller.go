package controllers

import (
    "hospital-api-go-pro/database"
    "hospital-api-go-pro/models"
    "net/http"

    "github.com/gin-gonic/gin"
)

func GetDoctors(c *gin.Context) {
    var doctors []models.Doctor
    database.DB.Find(&doctors)
    c.JSON(http.StatusOK, doctors)
}

func CreateDoctor(c *gin.Context) {
    var doctor models.Doctor
    if err := c.ShouldBindJSON(&doctor); err == nil {
        database.DB.Create(&doctor)
        c.JSON(http.StatusCreated, doctor)
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
}

func GetDoctorSchedule(c *gin.Context) {
    var doctors []models.Doctor
    database.DB.Select("name", "specialty", "schedule").Find(&doctors)
    c.JSON(http.StatusOK, doctors)
}
