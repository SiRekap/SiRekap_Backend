package server

import (
	"sirekap/SiRekap_Backend/controllers"
	"sirekap/SiRekap_Backend/middlewares"

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
	petugas := new(controllers.PetugasController)
	suaraC := new(controllers.SuaraCController)
	tps := new(controllers.TpsController)

	// Without token
	router.GET("/healthz", health.GetHealthStatus)
	router.POST("/register", petugas.RegisterPetugas)
	router.POST("/login", petugas.LoginPetugas)
	router.POST("/register-pemeriksa", petugas.RegisterPemeriksa)

	v1 := router.Group("v1")
	v1.Use(middlewares.Validate())

	// Form C Image
	v1.POST("/formc-image/send-formc-image-raw", formcImage.SendFormcImageRaw)
	v1.GET("/formc-image/get-formc-image-group-by-id-tps-and-jenis-pemilihan", formcImage.GetFormcImageGroupByIdTpsAndJenisPemilihan)

	// Form C Kesesuaian
	v1.POST("/formc-kesesuaian/send-formc-kesesuaian", formcKesesuaian.SendFormCKesesuaian)

	// Petugas
	v1.GET("/petugas/:id_petugas", petugas.GetPetugasTpsByIdPetugas)
	v1.GET("/petugas/all-pemeriksa", petugas.GetAllPemeriksaByTpsAndJenisPemilihan)

	// Suara C
	v1.GET("/suara-c/get-suara-c-final", suaraC.GetSuaraCFinal)
	v1.POST("/suara-c/send-suara-c-final", suaraC.SendSuaraCFinal)
	v1.GET("/suara-c/get-suara-c-proses", suaraC.GetSuaraCProses)
	v1.POST("/suara-c/send-suara-c-proses", suaraC.SendSuaraCProses)

	// TPS
	v1.GET("/tps/:id_tps", tps.GetTpsDetail)

	// External Component Test
	v1.POST("/kafka/send-results", controllers.SendFormcResultToKafkaTest)

	// router.Use(middlewares.AuthMiddleware())

	// v1 := router.Group("v1")
	// {
	// 	userGroup := v1.Group("user")
	// 	{
	// 		user := new(controllers.UserController)
	// 		userGroup.GET("/:id", user.Retrieve)
	// 	}
	// }

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
	// router.POST("/formc-image/send-formc-status-data", formcImage.SendFormcStatusData)
	// router.POST("/formc-image/send-formc-status-image", formcImage.SendFormcStatusImage)

	return router

}
