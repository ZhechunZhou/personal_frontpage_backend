package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// SetupRouter
/*
 Return: the root router of the web portal
*/
func SetupRouter() *gin.Engine {
	//create common router
	router := gin.Default()

	//configue
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://www.zhechundemo.com"},
		AllowMethods:     []string{"PUT, POST, GET, DELETE, OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//v1 := router.Group("/api/v1")

	//public  group

	return router
}
