package main

import (
	"os"
	"github.com/wkloose/Buddya/initializers"
	"github.com/wkloose/Buddya/middleware"
	//"github.com/wkloose/Buddya/Routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()

	router.Use(middleware.RequireAuth)
	RegisterRoutes(router)

	
	router.Run(":" + os.Getenv("PORT"))
}
