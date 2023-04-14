package models

import (
	"errors"
	"sirekap/SiRekap_Backend/db"

	"gorm.io/gorm"
)

type Wilayah struct {
	IdWilayah       int    `json:"id_wilayah" binding:"required" gorm:"primaryKey"`
	Tingkat         int    `json:"tingkat" binding:"required"`
	NamaWilayah     string `json:"nama_wilayah" binding:"required"`
	IdParentWilayah int    `json:"id_parent_wilayah" binding:"required"`
	KodeWilayah     string `json:"kode_wilayah" binding:"required"`
}

func (w Wilayah) GetWilayahById(idWilayah int) (Wilayah, error) {
	db := db.GetDB()

	wilayah := Wilayah{}

	result := db.Where("id_wilayah = ?", idWilayah).First(&wilayah)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return Wilayah{}, result.Error
	}

	return wilayah, nil
}
