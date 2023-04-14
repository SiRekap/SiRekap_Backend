package migrations

import (
	"sirekap/SiRekap_Backend/db"
	"sirekap/SiRekap_Backend/models"
)

func Migrate() {
	db := db.GetDB()

	db.AutoMigrate(&models.FormcImage{})
	db.AutoMigrate(&models.FormcImagePayload{})

	db.AutoMigrate(&models.FormcKesesuaian{})
	db.AutoMigrate(&models.Paslon{})

	db.AutoMigrate(&models.FormcAdministrasiHlmSatuProses{})
	db.AutoMigrate(&models.FormcAdministrasiHlmSatuFinal{})
	db.AutoMigrate(&models.FormcAdministrasiHlmDuaProses{})
	db.AutoMigrate(&models.FormcAdministrasiHlmDuaFinal{})

	db.AutoMigrate(&models.Paslon{})

	db.AutoMigrate(&models.PetugasTps{})
	db.AutoMigrate(&models.Pemeriksa{})

	db.AutoMigrate(&models.SuaraCFinal{})
	db.AutoMigrate(&models.SuaraCProses{})

	db.AutoMigrate(&models.Tps{})

	db.AutoMigrate(&models.Wilayah{})
}
