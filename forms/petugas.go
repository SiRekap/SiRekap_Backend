package forms

type PetugasRegisterData struct {
	Email    string `json:"email" binding:"required"`
	DeviceId string `json:"device_id" binding:"required"`
	Msisdn   string `json:"msisdn" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PemeriksaRegisterData struct {
	Nama           string `json:"nama" binding:"required"`
	Nik            string `json:"nik" binding:"required"`
	NoHandphone    string `json:"no_handphone" binding:"required"`
	JenisPemeriksa string `json:"jenis_pemeriksa" binding:"required"`
	JenisPemilihan string `json:"jenis_pemilihan" binding:"required"`
	IdPaslon       int    `json:"id_paslon" binding:"required"`
}
