package models

import (
	"errors"
	"sirekap/SiRekap_Backend/db"

	"gorm.io/gorm"
)

type Tps struct {
	IdTps          string `json:"id_tps" binding:"required" gorm:"primaryKey"`
	Alamat         string `json:"alamat" binding:"required"`
	Koordinat      string `json:"koordinat" binding:"required"`
	Tipe           string `json:"tipe" binding:"required"`
	IdWilayahDasar string `json:"id_wilayah_dasar" binding:"required"`
}

func (h Tps) GetById(idTps string) (Tps, error) {
	db := db.GetDB()

	tps := Tps{}

	result := db.Where("id_tps = ?", idTps).First(&tps)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return Tps{}, result.Error
	}

	return tps, nil
}
