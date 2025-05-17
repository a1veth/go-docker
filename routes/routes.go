package routes

import (
    "github.com/gin-gonic/gin"
    "go-api-template/controllers"
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/ping", controllers.Ping)
}
