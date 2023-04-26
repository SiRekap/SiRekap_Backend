package forms

type FullWilayahIdList struct {
	IdKelurahan     int `json:"id_kelurahan" binding:"required"`
	IdKecamatan     int `json:"id_kecamatan" binding:"required"`
	IdKabupatenKota int `json:"id_kabupaten_kota" binding:"required"`
	IdProvinsi      int `json:"id_provinsi" binding:"required"`
}
