package models

import (
	"sirekap/SiRekap_Backend/db"
)

type (
	FormcImagePayload struct {
		IdImage int    `json:"id_image" binding:"required" gorm:"primaryKey"`
		Payload string `json:"payload" binding:"required"`
	}

	FormcImage struct {
		IdImage        int    `json:"id_image" binding:"required" gorm:"primaryKey"`
		JenisImage     int    `json:"jenis_image" binding:"required"`
		NoLembar       int    `json:"no_lembar" binding:"required"`
		IdTps          int    `json:"id_tps" binding:"required"`
		JenisPemilihan string `json:"jenis_pemilihan" binding:"required"`
		Uuid           string `json:"uuid" binding:"required"`
		NamaFile       string `json:"nama_file" binding:"required"`
		FileHash       string `json:"file_hash" binding:"required"`
		FileSignature  string `json:"file_signature" binding:"required"`
		KodeUsl        string `json:"kode_usl" binding:"required"`
	}

	FormcStatusData struct {
		IdImage     int    `json:"id_image" binding:"required" gorm:"primaryKey"`
		Status      string `json:"status" binding:"required"`
		IdVersiMaks int    `json:"id_versi_maks" binding:"required"`
		Keterangan  string `json:"keterangan" binding:"required"`
	}

	FormcStatusImage struct {
		IdImage    int    `json:"id_image" binding:"required" gorm:"primaryKey"`
		Status     string `json:"status" binding:"required"`
		Keterangan string `json:"keterangan" binding:"required"`
	}
)

func (f FormcImagePayload) SendFormcImagePayload(formImagePayload FormcImagePayload) (*FormcImagePayload, error) {
	db := db.GetDB()

	db.Create(&formImagePayload)

	return &formImagePayload, nil
}

func (f FormcImage) SendFormcImage(formImage FormcImage) (*FormcImage, error) {
	db := db.GetDB()

	db.Create(&formImage)

	return &formImage, nil
}

func (f FormcStatusData) SendFormcStatusData(formStatusData FormcStatusData) (*FormcStatusData, error) {
	db := db.GetDB()

	db.Create(&formStatusData)

	return &formStatusData, nil
}

func (f FormcStatusImage) SendFormcStatusImage(formStatusImage FormcStatusImage) (*FormcStatusImage, error) {
	db := db.GetDB()

	db.Create(&formStatusImage)

	return &formStatusImage, nil
}
