package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wkloose/Buddya/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/budaya", controllers.GetBudayaByCity)
	}
}
