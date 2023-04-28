package models

import (
	"errors"
	"sirekap/SiRekap_Backend/db"
	"sirekap/SiRekap_Backend/forms"

	"gorm.io/gorm"
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
		JenisPemilihan int    `json:"jenis_pemilihan" binding:"required"`
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

	FormcImageGroup struct {
		IdGroup        int    `json:"id_group" binding:"required" gorm:"primaryKey"`
		IdTps          int    `json:"id_tps" binding:"required"`
		JenisPemilihan int    `json:"jenis_pemilihan" binding:"required"`
		IdImageHlm1    int    `json:"id_image_hlm1" binding:"required"`
		IdImageHlm2    int    `json:"id_image_hlm2" binding:"required"`
		IdImageHlm3    int    `json:"id_image_hlm3" binding:"required"`
		PdfUrl         string `json:"pdf_url" binding:"required"`
	}
)

func (f FormcImagePayload) SendFormcImagePayload() (*FormcImagePayload, error) {
	db := db.GetDB()

	db.Create(&f)

	return &f, nil
}

func (f FormcImage) GetFormcImage(idImage int) error {
	db := db.GetDB()

	f = FormcImage{
		IdImage: idImage,
	}

	result := db.First(&f)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	return nil
}

func (f FormcImage) SendFormcImage() (*FormcImage, error) {
	db := db.GetDB()

	db.Create(&f)

	return &f, nil
}

func (f FormcStatusData) SendFormcStatusData() (*FormcStatusData, error) {
	db := db.GetDB()

	db.Create(&f)

	return &f, nil
}

func (f FormcStatusImage) SendFormcStatusImage() (*FormcStatusImage, error) {
	db := db.GetDB()

	db.Create(&f)

	return &f, nil
}

func (f FormcImageGroup) SendFormcImageGroup() (*FormcImageGroup, error) {
	db := db.GetDB()

	db.Create(&f)

	return &f, nil
}

func SendFormcImageRaw(formcImageRaw forms.FormcImageRaw) (*forms.FormcImageRawResponse, error) {

	idImageList := make([]int, 0)

	for i := 0; i < 3; i++ {

		formImagePayload := FormcImagePayload{
			Payload: formcImageRaw.PayloadList[i],
		}

		resPayload, _ := formImagePayload.SendFormcImagePayload()

		formcImage := FormcImage{
			IdImage:        resPayload.IdImage,
			NoLembar:       i + 1,
			IdTps:          formcImageRaw.IdTps,
			JenisPemilihan: formcImageRaw.JenisPemilihan,
		}

		formcImage.SendFormcImage()

		idImageList = append(idImageList, resPayload.IdImage)
	}

	return &forms.FormcImageRawResponse{
		FormcImageRaw: formcImageRaw,
		IdImageList:   idImageList,
	}, nil
}
