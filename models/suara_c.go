package models

import (
	"errors"
	"sirekap/SiRekap_Backend/db"

	"gorm.io/gorm"
)

type (
	SuaraCFinal struct {
		IdVersi    int    `json:"id_versi" binding:"required" gorm:"primaryKey"`
		IdImage    int    `json:"id_image" binding:"required"`
		IdPaslon   int    `json:"id_paslon" binding:"required"`
		JmlSuara   int    `json:"jml_suara" binding:"required"`
		Keterangan string `json:"keterangan" binding:"required"`
	}

	SuaraCProses struct {
		IdImage     int    `json:"id_image" binding:"required" gorm:"primaryKey"`
		IdPaslon    int    `json:"id_paslon" binding:"required" gorm:"primaryKey"`
		JmlSuaraOcr int    `json:"jml_suara_ocr" binding:"required"`
		JmlSuaraOmr int    `json:"jml_suara_omr" binding:"required"`
		Keterangan  string `json:"keterangan" binding:"required"`
	}
)

func (s SuaraCFinal) GetSuaraCFinal(idVersi int) (SuaraCFinal, error) {
	db := db.GetDB()

	suara := SuaraCFinal{
		IdVersi: idVersi,
	}

	result := db.First(&suara)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return SuaraCFinal{}, result.Error
	}

	return suara, nil
}

func (s SuaraCFinal) SendSuaraCFinal(suara SuaraCFinal) (*SuaraCFinal, error) {
	db := db.GetDB()

	db.Create(&suara)

	return &suara, nil
}

func (s SuaraCProses) GetSuaraCProses(idImage int) (SuaraCProses, error) {
	db := db.GetDB()

	suara := SuaraCProses{
		IdImage: idImage,
	}

	result := db.First(&suara)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return SuaraCProses{}, result.Error
	}

	return suara, nil
}

func (s SuaraCProses) SendSuaraCProses(suara SuaraCProses) (*SuaraCProses, error) {
	db := db.GetDB()

	db.Create(&suara)

	return &suara, nil
}
