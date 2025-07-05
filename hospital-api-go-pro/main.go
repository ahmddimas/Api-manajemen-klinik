package main

import (
    "hospital-api-go-pro/database"
    "hospital-api-go-pro/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    database.ConnectDB()
    routes.SetupRoutes(r)
    r.Run(":8080")
}
