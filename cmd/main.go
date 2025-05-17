package main

import (
    "github.com/gin-gonic/gin"
    "go-api-template/routes"
)

func main() {
    r := gin.Default()
    routes.SetupRoutes(r)
    r.Run(":8080")
}
