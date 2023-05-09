package models

import "sirekap/SiRekap_Backend/db"

type (
	FormcResult struct {
		IdFormcResult   int `json:"id_image" binding:"required" gorm:"primaryKey"`
		IdTps           int `json:"id_tps" binding:"required"`
		IdKelurahan     int `json:"id_kelurahan" binding:"required"`
		IdKecamatan     int `json:"id_kecamatan" binding:"required"`
		IdKabupatenKota int `json:"id_kabupaten_kota" binding:"required"`
		IdProvinsi      int `json:"id_provinsi" binding:"required"`
		IdPaslon        int `json:"id_paslon" binding:"required"`
		JmlSuara        int `json:"jml_suara" binding:"required"`
		JenisPemilihan  int `json:"jenis_pemilihan" binding:"required"`
	}
)

func SendFormcResult(form FormcResult) (*FormcResult, error) {
	db := db.GetDB()

	db.Create(&form)

	return &form, nil
}
