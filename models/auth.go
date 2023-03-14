package models

import (
	"errors"
	"sirekap/SiRekap_Backend/db"
	"sirekap/SiRekap_Backend/forms"

	"gorm.io/gorm"
)

type PetugasTps struct {
	IdPetugas      string `json:"id_petugas" binding:"required" gorm:"primaryKey"`
	IdWilayah      string `json:"id_wilayah" binding:"required"`
	JenisPemeriksa string `json:"jenis_pemeriksa" binding:"required"`
	IdPaslon       string `json:"id_paslon" binding:"required"`
	Urutan         string `json:"urutan" binding:"required"`
	Nama           string `json:"nama" binding:"required"`
	NoHandphone    string `json:"no_handphone" binding:"required"`
	Nik            string `json:"nik" binding:"required"`
	Email          string `json:"email" binding:"required"`
	Msisdn         string `json:"msisdn" binding:"required"`
}

func (h PetugasTps) Register(userRegisterData forms.UserRegisterData) (*PetugasTps, error) {
	db := db.GetDB()

	petugasTps, err := h.GetByEmail(userRegisterData.Email)
	if err != nil {
		return nil, err
	}

	db.Model(&petugasTps).Update("msisdn", userRegisterData.Msisdn)

	return &petugasTps, nil
}

func (h PetugasTps) GetByEmail(email string) (PetugasTps, error) {
	db := db.GetDB()

	petugasTps := PetugasTps{}

	result := db.Where("email = ?", email).First(&petugasTps)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return PetugasTps{}, result.Error
	}

	return petugasTps, nil
}
