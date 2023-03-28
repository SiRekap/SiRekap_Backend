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
	// auth := new(controllers.AuthController)
	tps := new(controllers.TpsController)
	petugas := new(controllers.PetugasController)
	formc := new(controllers.FormCController)

	// Health
	router.GET("/healthz", health.GetHealthStatus)

	// TPS
	router.GET("/tps/:id_tps", tps.GetTpsDetail)

	// Petugas
	router.POST("/register", petugas.RegisterPetugas)
	router.POST("/register-pemeriksa", petugas.RegisterPemeriksa)

	// Form C
	router.GET("/formc/get-formc-administrasi-hlm-satu-proses", formc.GetFormCAdministrasiHlmSatuProses)
	router.GET("/formc/get-formc-administrasi-hlm-satu-final", formc.GetFormCAdministrasiHlmSatuFinal)
	router.GET("/formc/get-formc-administrasi-hlm-satu", formc.GetFormCAdministrasiHlmSatu)

	router.POST("/formc/send-formc-administrasi-hlm-satu-proses", formc.SendFormCAdministrasiHlmSatuProses)
	router.POST("/formc/send-formc-administrasi-hlm-satu-final", formc.SendFormCAdministrasiHlmSatuFinal)
	router.POST("/formc/send-formc-administrasi-hlm-satu", formc.SendFormCAdministrasiHlmSatu)

	router.GET("/formc/get-formc-administrasi-hlm-dua-proses", formc.GetFormCAdministrasiHlmDuaProses)
	router.GET("/formc/get-formc-administrasi-hlm-dua-final", formc.GetFormCAdministrasiHlmDuaFinal)
	router.GET("/formc/get-formc-administrasi-hlm-dua", formc.GetFormCAdministrasiHlmDua)

	router.POST("/formc/send-formc-administrasi-hlm-dua-proses", formc.SendFormCAdministrasiHlmDuaProses)
	router.POST("/formc/send-formc-administrasi-hlm-dua-final", formc.SendFormCAdministrasiHlmDuaFinal)
	router.POST("/formc/send-formc-administrasi-hlm-dua", formc.SendFormCAdministrasiHlmDua)

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
