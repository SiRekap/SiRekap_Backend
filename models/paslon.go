package models

import (
	"errors"
	"sirekap/SiRekap_Backend/db"

	"gorm.io/gorm"
)

type Paslon struct {
	IdPaslon         int    `json:"id_paslon" binding:"required" gorm:"primaryKey"`
	IdJnsPencalonan  string `json:"id_jns_pencalonan" binding:"required"`
	JenisPemilihan   string `json:"jenis_pemilihan" binding:"required"`
	IdWilayah        int    `json:"id_wilayah" binding:"required"`
	NoUrutPencalonan int    `json:"no_urut_pencalonan" binding:"required"`
	NamaCaka         string `json:"nama_caka" binding:"required"`
	NamaCawaka       string `json:"nama_cawaka" binding:"required"`
}

func (p Paslon) GetPaslonById(idPaslon int) (Paslon, error) {
	db := db.GetDB()

	paslon := Paslon{
		IdPaslon: idPaslon,
	}

	result := db.First(&paslon)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return Paslon{}, result.Error
	}

	return paslon, nil
}
