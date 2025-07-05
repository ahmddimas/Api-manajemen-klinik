package database

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
    "hospital-api-go-pro/models"
)

var DB *gorm.DB

func ConnectDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("hospital.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Gagal koneksi ke database:", err)
    }

    DB.AutoMigrate(&models.Patient{}, &models.Doctor{})
}
