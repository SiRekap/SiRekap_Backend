package models

import (
	"errors"
	"sirekap/SiRekap_Backend/db"
	"sirekap/SiRekap_Backend/forms"

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

func GetFullWilayahIdList(idWilayahKelurahan int) (forms.FullWilayahIdList, error) {
	db := db.GetDB()

	kelurahan := Wilayah{}

	result := db.Where("id_wilayah = ?", idWilayahKelurahan).First(&kelurahan)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return forms.FullWilayahIdList{}, result.Error
	}

	kecamatan := Wilayah{}

	result = db.Where("id_wilayah = ?", kelurahan.IdParentWilayah).First(&kecamatan)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return forms.FullWilayahIdList{}, result.Error
	}

	kabupatenKota := Wilayah{}

	result = db.Where("id_wilayah = ?", kecamatan.IdParentWilayah).First(&kabupatenKota)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return forms.FullWilayahIdList{}, result.Error
	}

	provinsi := Wilayah{}

	result = db.Where("id_wilayah = ?", kabupatenKota.IdParentWilayah).First(&provinsi)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return forms.FullWilayahIdList{}, result.Error
	}

	return forms.FullWilayahIdList{
		IdKelurahan:     kelurahan.IdWilayah,
		IdKecamatan:     kecamatan.IdWilayah,
		IdKabupatenKota: kabupatenKota.IdWilayah,
		IdProvinsi:      provinsi.IdWilayah,
	}, nil
}
