package server

import (
	"sirekap/SiRekap_Backend/controllers"

	"github.com/gin-gonic/gin"
	// "sirekap/SiRekap_Backend/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)
	auth := new(controllers.AuthController)

	router.GET("/health", health.Status)

	router.POST("/register", auth.Register)
	// router.Use(middlewares.AuthMiddleware())

	// v1 := router.Group("v1")
	// {
	// 	userGroup := v1.Group("user")
	// 	{
	// 		user := new(controllers.UserController)
	// 		userGroup.GET("/:id", user.Retrieve)
	// 	}
	// }
	return router

}
