package main

import (
	"os"
	"github.com/wkloose/Buddya/initializers"
	"github.com/gin-gonic/gin"
	"github.com/wkloose/Buddya/Routes"
	"github.com/wkloose/Buddya/Seeder"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.ResetDatabase()
}

func main() {
	Seeder.SeedAll()
	router := gin.Default()

	Routes.RegisterRoutes(router)

	
	router.Run(":" + os.Getenv("PORT"))
}
