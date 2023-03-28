package models

import (
	"errors"
	"sirekap/SiRekap_Backend/db"
	"sirekap/SiRekap_Backend/forms"
	"time"

	"gorm.io/gorm"
)

type (
	PetugasTps struct {
		IdPetugas      int    `json:"id_petugas" binding:"required" gorm:"primaryKey"`
		IdWilayah      int    `json:"id_wilayah" binding:"required"`
		JenisPemeriksa string `json:"jenis_pemeriksa" binding:"required"`
		IdPaslon       int    `json:"id_paslon" binding:"required"`
		Urutan         string `json:"urutan" binding:"required"`
		Nama           string `json:"nama" binding:"required"`
		NoHandphone    string `json:"no_handphone" binding:"required"`
		Nik            string `json:"nik" binding:"required"`
		Email          string `json:"email" binding:"required"`
		Msisdn         string `json:"msisdn" binding:"required"`
	}

	Pemeriksa struct {
		IdPemeriksa     int        `json:"id_pemeriksa" binding:"required" gorm:"primaryKey"`
		JenisPemilihan  string     `json:"jenis_pemilihan" binding:"required"`
		Url             string     `json:"url" binding:"required"`
		WaktuUrl        *time.Time `json:"waktu_url" binding:"required"`
		KeteranganUrl   string     `json:"keterangan_url" binding:"required"`
		HashUrl         string     `json:"hash_url" binding:"required"`
		WaktuAkses      *time.Time `json:"waktu_akses" binding:"required"`
		KeteranganAkses string     `json:"keterangan_akses" binding:"required"`
		HashAkses       string     `json:"hash_akses" binding:"required"`
	}
)

func (p PetugasTps) RegisterPetugas(userRegisterData forms.PetugasRegisterData) (*PetugasTps, error) {
	db := db.GetDB()

	petugasTps, err := p.GetByEmail(userRegisterData.Email)
	if err != nil {
		return nil, err
	}

	db.Model(&petugasTps).Update("msisdn", userRegisterData.Msisdn)

	return &petugasTps, nil
}

func (p PetugasTps) GetByEmail(email string) (PetugasTps, error) {
	db := db.GetDB()

	petugasTps := PetugasTps{}

	result := db.Where("email = ?", email).First(&petugasTps)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return PetugasTps{}, result.Error
	}

	return petugasTps, nil
}

func (p Pemeriksa) RegisterPemeriksa(userRegisterData forms.PemeriksaRegisterData) (*PetugasTps, error) {
	db := db.GetDB()

	petugasTps := PetugasTps{
		JenisPemeriksa: userRegisterData.JenisPemeriksa,
		IdPaslon:       userRegisterData.IdPaslon,
		Nama:           userRegisterData.Nama,
		Nik:            userRegisterData.Nik,
		NoHandphone:    userRegisterData.NoHandphone,
	}

	db.Create(&petugasTps)

	pemeriksa := Pemeriksa{
		IdPemeriksa:    petugasTps.IdPetugas,
		JenisPemilihan: userRegisterData.JenisPemilihan,
	}

	db.Create(&pemeriksa)

	return &petugasTps, nil
}
