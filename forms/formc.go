package forms

type FormcImageRaw struct {
	NoLembar       int    `json:"no_lembar" binding:"required"`
	IdTps          int    `json:"id_tps" binding:"required"`
	JenisPemilihan string `json:"jenis_pemilihan" binding:"required"`
	Payload        string `json:"payload" binding:"required"`
}
