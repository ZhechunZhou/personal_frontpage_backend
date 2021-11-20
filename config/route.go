package config

import (
	"git.wisecar.co/wisecar/web-api-golang/claim"
	"git.wisecar.co/wisecar/web-api-golang/countrystate"
	"git.wisecar.co/wisecar/web-api-golang/customerInfo"
	"git.wisecar.co/wisecar/web-api-golang/driver"
	"git.wisecar.co/wisecar/web-api-golang/insurance"
	"git.wisecar.co/wisecar/web-api-golang/registration"
	"git.wisecar.co/wisecar/web-api-golang/service"
	"git.wisecar.co/wisecar/web-api-golang/share"
	"git.wisecar.co/wisecar/web-api-golang/user"
	"git.wisecar.co/wisecar/web-api-golang/vehicle"
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
		AllowOrigins:     []string{"http://localhost:3000", "https://portal.wisecar.co"},
		AllowMethods:     []string{"PUT, POST, GET, DELETE, OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v1 := router.Group("/api/v1")

	//customerInfo login routing group
	customerInfo.PublicApi(v1)
	countrystate.Routing(v1)
	authenticated := router.Group("/api/v1")
	authenticated.Use(customerInfo.AuthMiddleware(true))
	{
		customerInfo.Routing(authenticated)
		user.WebDetail(authenticated)
		vehicle.WebDetails(authenticated)
		insurance.WebDetail(authenticated)
		registration.WebDetail(authenticated)
		driver.WebDetail(authenticated)
		claim.WebDetail(authenticated)
		service.WebDetail(authenticated)
		share.WebDetails(authenticated)
	}
	return router
}
