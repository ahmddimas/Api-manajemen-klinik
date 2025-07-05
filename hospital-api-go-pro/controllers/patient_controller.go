package controllers

import (
    "hospital-api-go-pro/database"
    "hospital-api-go-pro/models"
    "net/http"

    "github.com/gin-gonic/gin"
)

func GetPatients(c *gin.Context) {
    var patients []models.Patient
    database.DB.Find(&patients)
    c.JSON(http.StatusOK, patients)
}

func CreatePatient(c *gin.Context) {
    var patient models.Patient
    if err := c.ShouldBindJSON(&patient); err == nil {
        database.DB.Create(&patient)
        c.JSON(http.StatusCreated, patient)
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
}
