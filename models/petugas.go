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
		IdTps          int    `json:"id_wilayah" binding:"required"`
		JenisPemilihan int    `json:"jenis_pemilihan" binding:"required"`
		JenisPemeriksa int    `json:"jenis_pemeriksa" binding:"required"`
		IdPaslon       int    `json:"id_paslon" binding:"required"`
		Urutan         string `json:"urutan" binding:"required"`
		Nama           string `json:"nama" binding:"required"`
		NoHandphone    string `json:"no_handphone" binding:"required"`
		Nik            string `json:"nik" binding:"required"`
		Email          string `json:"email" binding:"required"`
		Msisdn         string `json:"msisdn" binding:"required"`
		DeviceId       string `json:"device_id" binding:"required"`
		Password       string `json:"password" binding:"required"`
	}

	Pemeriksa struct {
		IdPemeriksa     int        `json:"id_pemeriksa" binding:"required" gorm:"primaryKey"`
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

	db.Model(&petugasTps).Updates(PetugasTps{
		Msisdn:   userRegisterData.Msisdn,
		DeviceId: userRegisterData.DeviceId,
		Password: userRegisterData.Password,
	})

	return &petugasTps, nil
}

func (p PetugasTps) LoginPetugas(loginData forms.PetugasLoginData) (PetugasTps, error) {
	petugasTps, err := p.GetByEmailAndPassword(loginData.Email, loginData.Password)
	if err != nil {
		return PetugasTps{}, err
	}

	return petugasTps, nil
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

func (p PetugasTps) GetByEmailAndPassword(email string, password string) (PetugasTps, error) {
	db := db.GetDB()

	petugasTps := PetugasTps{}

	result := db.Where("email = ? AND password = ?", email, password).First(&petugasTps)

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
		JenisPemilihan: userRegisterData.JenisPemilihan,
	}

	db.Create(&petugasTps)

	pemeriksa := Pemeriksa{
		IdPemeriksa: petugasTps.IdPetugas,
	}

	db.Create(&pemeriksa)

	return &petugasTps, nil
}

func (p Pemeriksa) GetPetugasTpsByIdPetugas(idPetugas int) (PetugasTps, error) {
	db := db.GetDB()

	petugasTps := PetugasTps{}

	result := db.Where("id_petugas = ?", idPetugas).First(&petugasTps)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return PetugasTps{}, result.Error
	}

	return petugasTps, nil
}

func (p Pemeriksa) GetAllPemeriksaByTpsAndJenisPemilihan(idTps int, jenisPemilihan int) ([]PetugasTps, error) {
	db := db.GetDB()

	petugasTpsList := []PetugasTps{}

	result := db.Where("id_tps = ? AND jenis_pemilihan = ?", idTps, jenisPemilihan).
		Where("jenis_pemeriksa IS NOT NULL").
		Find(&petugasTpsList)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		print(result.Error)
		return []PetugasTps{}, result.Error
	}

	return petugasTpsList, nil
}
