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

	// formcAdministrasi := new(controllers.FormcAdministrasiController)
	formcImage := new(controllers.FormcImageController)
	formcKesesuaian := new(controllers.FormcKesesuaianController)
	health := new(controllers.HealthController)
	// auth := new(controllers.AuthController)
	petugas := new(controllers.PetugasController)
	suaraC := new(controllers.SuaraCController)
	tps := new(controllers.TpsController)

	// Form C Administrasi
	// router.GET("/formc-administrasi/get-formc-administrasi-hlm-satu-proses", formcAdministrasi.GetFormcAdministrasiHlmSatuProses)
	// router.GET("/formc-administrasi/get-formc-administrasi-hlm-satu-final", formcAdministrasi.GetFormcAdministrasiHlmSatuFinal)
	// router.GET("/formc-administrasi/get-formc-administrasi-hlm-satu", formcAdministrasi.GetFormcAdministrasiHlmSatu)

	// router.POST("/formc-administrasi/send-formc-administrasi-hlm-satu-proses", formcAdministrasi.SendFormcAdministrasiHlmSatuProses)
	// router.POST("/formc-administrasi/send-formc-administrasi-hlm-satu-final", formcAdministrasi.SendFormcAdministrasiHlmSatuFinal)
	// router.POST("/formc-administrasi/send-formc-administrasi-hlm-satu", formcAdministrasi.SendFormcAdministrasiHlmSatu)

	// router.GET("/formc-administrasi/get-formc-administrasi-hlm-dua-proses", formcAdministrasi.GetFormcAdministrasiHlmDuaProses)
	// router.GET("/formc-administrasi/get-formc-administrasi-hlm-dua-final", formcAdministrasi.GetFormcAdministrasiHlmDuaFinal)
	// router.GET("/formc/get-formc-administrasi-hlm-dua", formcAdministrasi.GetFormcAdministrasiHlmDua)

	// router.POST("/formc-administrasi/send-formc-administrasi-hlm-dua-proses", formcAdministrasi.SendFormcAdministrasiHlmDuaProses)
	// router.POST("/formc-administrasi/send-formc-administrasi-hlm-dua-final", formcAdministrasi.SendFormcAdministrasiHlmDuaFinal)
	// router.POST("/formc/send-formc-administrasi-hlm-dua", formcAdministrasi.SendFormcAdministrasiHlmDua)

	// Form C Image
	// router.POST("/formc-image/send-formc-image-payload", formcImage.SendFormcImagePayload)
	// router.POST("/formc-image/send-formc-image", formcImage.SendFormcImage)
	router.POST("/formc-image/send-formc-image-raw", formcImage.SendFormcImageRaw)
	// router.POST("/formc-image/send-formc-status-data", formcImage.SendFormcStatusData)
	// router.POST("/formc-image/send-formc-status-image", formcImage.SendFormcStatusImage)
	router.GET("/formc-image/get-formc-image-group-by-id-tps-and-jenis-pemilihan", formcImage.GetFormcImageGroupByIdTpsAndJenisPemilihan)

	// Form C Kesesuaian
	router.POST("/formc-kesesuaian/send-formc-kesesuaian", formcKesesuaian.SendFormCKesesuaian)

	// Health
	router.GET("/healthz", health.GetHealthStatus)

	// Petugas
	router.POST("/register", petugas.RegisterPetugas)
	router.POST("/register-pemeriksa", petugas.RegisterPemeriksa)
	router.GET("/petugas/:id_petugas", petugas.GetPetugasTpsByIdPetugas)
	router.GET("/petugas/all-pemeriksa", petugas.GetAllPemeriksaByTpsAndJenisPemilihan)

	// Suara C
	router.GET("/suara-c/get-suara-c-final", suaraC.GetSuaraCFinal)
	router.POST("/suara-c/send-suara-c-final", suaraC.SendSuaraCFinal)
	router.GET("/suara-c/get-suara-c-proses", suaraC.GetSuaraCProses)
	router.POST("/suara-c/send-suara-c-proses", suaraC.SendSuaraCProses)

	// TPS
	router.GET("/tps/:id_tps", tps.GetTpsDetail)

	// External Component Test
	router.POST("/kafka/send-results", controllers.SendFormcResultToKafkaTest)

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
