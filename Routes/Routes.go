package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wkloose/Buddya/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/budaya", controllers.GetBudayaByCity)
		api.GET("/wisata", controllers.GetWisataByCity)
		api.GET("/kuliner", controllers.GetKulinerByCity)
		api.GET("/events", controllers.GetEventByCity)
		api.GET("/places", controllers.GetPlaces)
	}
}
