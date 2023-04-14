package models

import (
	"sirekap/SiRekap_Backend/db"
)

type (
	FormcKesesuaian struct {
		IdImage         int    `json:"id_image" binding:"required" gorm:"primaryKey"`
		IdPemeriksa     string `json:"id_pemeriksa" binding:"required" gorm:"primaryKey"`
		IsSesuai        bool   `json:"is_sesuai" binding:"required"`
		Komentar        string `json:"komentar" binding:"required"`
		IsSesuaiPerItem bool   `json:"is_sesuai_per_item" binding:"required"`
		KoreksiPerItem  []int  `json:"koreksi_per_item" binding:"required"`
	}
)

func (f FormcKesesuaian) SendFormCKesesuaian(form FormcKesesuaian) (*FormcKesesuaian, error) {
	db := db.GetDB()

	db.Create(&form)

	return &form, nil
}
