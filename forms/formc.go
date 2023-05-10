package forms

type FormcImageRaw struct {
	IdTps          int      `json:"id_tps" binding:"required"`
	JenisPemilihan int      `json:"jenis_pemilihan" binding:"required"`
	PayloadList    []string `json:"payload_list" binding:"required"`
	IdPaslonList   []int    `json:"id_paslon_list" binding:"required"`
}

type FormcImageRawResponse struct {
	FormcImageRaw FormcImageRaw
	IdImageList   []int
}

type FormcImageVisionRequest struct {
	IdImageList  []int    `json:"IdImageList" binding:"required"`
	ImageUrlList []string `json:"ImageUrlList" binding:"required"`
	IdPaslonList []int    `json:"IdPaslonList" binding:"required"`
}

type FormcImageVisionResponse struct {
	PemilihDptL   int `json:"pemilih_dpt_l" binding:"required"`
	PemilihDptP   int `json:"pemilih_dpt_p" binding:"required"`
	PemilihDptJ   int `json:"pemilih_dpt_j" binding:"required"`
	PemilihDpphL  int `json:"pemilih_dpph_l" binding:"required"`
	PemilihDpphP  int `json:"pemilih_dpph_p" binding:"required"`
	PemilihDpphJ  int `json:"pemilih_dpph_j" binding:"required"`
	PemilihDptbL  int `json:"pemilih_dptb_l" binding:"required"`
	PemilihDptbP  int `json:"pemilih_dptb_p" binding:"required"`
	PemilihDptbJ  int `json:"pemilih_dptb_j" binding:"required"`
	PemilihTotalL int `json:"pemilih_total_l" binding:"required"`
	PemilihTotalP int `json:"pemilih_total_p" binding:"required"`
	PemilihTotalJ int `json:"pemilih_total_j" binding:"required"`

	PenggunaDptL   int `json:"pengguna_dpt_l" binding:"required"`
	PenggunaDptP   int `json:"pengguna_dpt_p" binding:"required"`
	PenggunaDptJ   int `json:"pengguna_dpt_j" binding:"required"`
	PenggunaDpphL  int `json:"pengguna_dpph_l" binding:"required"`
	PenggunaDpphP  int `json:"pengguna_dpph_p" binding:"required"`
	PenggunaDpphJ  int `json:"pengguna_dpph_j" binding:"required"`
	PenggunaDptbL  int `json:"pengguna_dptb_l" binding:"required"`
	PenggunaDptbP  int `json:"pengguna_dptb_p" binding:"required"`
	PenggunaDptbJ  int `json:"pengguna_dptb_j" binding:"required"`
	PenggunaTotalL int `json:"pengguna_total_l" binding:"required"`
	PenggunaTotalP int `json:"pengguna_total_p" binding:"required"`
	PenggunaTotalJ int `json:"pengguna_total_j" binding:"required"`

	PemilihDisabilitasL  int `json:"pemilih_disabilitas_l" binding:"required"`
	PemilihDisabilitasP  int `json:"pemilih_disabilitas_p" binding:"required"`
	PemilihDisabilitasJ  int `json:"pemilih_disabilitas_j" binding:"required"`
	PenggunaDisabilitasL int `json:"pengguna_disabilitas_l" binding:"required"`
	PenggunaDisabilitasP int `json:"pengguna_disabilitas_p" binding:"required"`
	PenggunaDisabilitasJ int `json:"pengguna_disabilitas_j" binding:"required"`

	SuratDiterima       int `json:"surat_diterima" binding:"required"`
	SuratDikembalikan   int `json:"surat_dikembalikan" binding:"required"`
	SuratTidakDigunakan int `json:"surat_tidak_dikembalikan" binding:"required"`
	SuratDigunakan      int `json:"surat_digunakan" binding:"required"`

	SuaraSah            int `json:"suara_sah" binding:"required"`
	SuaraTidakSah       int `json:"suara_tidak_sah" binding:"required"`
	SuaraTotal          int `json:"suara_total" binding:"required"`
	PenggunaHakPilih    int `json:"pengguna_hak_pilih" binding:"required"`
	SuratSuaraDigunakan int `json:"surat_suara_digunakan" binding:"required"`

	JmlSuaraOcrList []int `json:"jml_suara_ocr_list" binding:"required"`
	JmlSuaraOmrList []int `json:"jml_suara_omr_list" binding:"required"`

	IdImageList  []int `json:"IdImageList" binding:"required"`
	IdPaslonList []int `json:"IdPaslonList" binding:"required"`
}

type FormcResultStreamProcessingRequest struct {
	IdTps           int `json:"id_tps" binding:"required"`
	IdKelurahan     int `json:"id_kelurahan" binding:"required"`
	IdKecamatan     int `json:"id_kecamatan" binding:"required"`
	IdKabupatenKota int `json:"id_kabupaten_kota" binding:"required"`
	IdProvinsi      int `json:"id_provinsi" binding:"required"`
	IdPaslon        int `json:"id_paslon" binding:"required"`
	JmlSuara        int `json:"jml_suara" binding:"required"`
	JenisPemilihan  int `json:"jenis_pemilihan" binding:"required"`
}

type FormcScanResponse struct {
	PemilihDptL   int `json:"pemilih_dpt_l" binding:"required"`
	PemilihDptP   int `json:"pemilih_dpt_p" binding:"required"`
	PemilihDptJ   int `json:"pemilih_dpt_j" binding:"required"`
	PemilihDpphL  int `json:"pemilih_dpph_l" binding:"required"`
	PemilihDpphP  int `json:"pemilih_dpph_p" binding:"required"`
	PemilihDpphJ  int `json:"pemilih_dpph_j" binding:"required"`
	PemilihDptbL  int `json:"pemilih_dptb_l" binding:"required"`
	PemilihDptbP  int `json:"pemilih_dptb_p" binding:"required"`
	PemilihDptbJ  int `json:"pemilih_dptb_j" binding:"required"`
	PemilihTotalL int `json:"pemilih_total_l" binding:"required"`
	PemilihTotalP int `json:"pemilih_total_p" binding:"required"`
	PemilihTotalJ int `json:"pemilih_total_j" binding:"required"`

	PenggunaDptL   int `json:"pengguna_dpt_l" binding:"required"`
	PenggunaDptP   int `json:"pengguna_dpt_p" binding:"required"`
	PenggunaDptJ   int `json:"pengguna_dpt_j" binding:"required"`
	PenggunaDpphL  int `json:"pengguna_dpph_l" binding:"required"`
	PenggunaDpphP  int `json:"pengguna_dpph_p" binding:"required"`
	PenggunaDpphJ  int `json:"pengguna_dpph_j" binding:"required"`
	PenggunaDptbL  int `json:"pengguna_dptb_l" binding:"required"`
	PenggunaDptbP  int `json:"pengguna_dptb_p" binding:"required"`
	PenggunaDptbJ  int `json:"pengguna_dptb_j" binding:"required"`
	PenggunaTotalL int `json:"pengguna_total_l" binding:"required"`
	PenggunaTotalP int `json:"pengguna_total_p" binding:"required"`
	PenggunaTotalJ int `json:"pengguna_total_j" binding:"required"`

	PemilihDisabilitasL  int `json:"pemilih_disabilitas_l" binding:"required"`
	PemilihDisabilitasP  int `json:"pemilih_disabilitas_p" binding:"required"`
	PemilihDisabilitasJ  int `json:"pemilih_disabilitas_j" binding:"required"`
	PenggunaDisabilitasL int `json:"pengguna_disabilitas_l" binding:"required"`
	PenggunaDisabilitasP int `json:"pengguna_disabilitas_p" binding:"required"`
	PenggunaDisabilitasJ int `json:"pengguna_disabilitas_j" binding:"required"`

	SuratDiterima       int `json:"surat_diterima" binding:"required"`
	SuratDikembalikan   int `json:"surat_dikembalikan" binding:"required"`
	SuratTidakDigunakan int `json:"surat_tidak_dikembalikan" binding:"required"`
	SuratDigunakan      int `json:"surat_digunakan" binding:"required"`

	SuaraSah            int `json:"suara_sah" binding:"required"`
	SuaraTidakSah       int `json:"suara_tidak_sah" binding:"required"`
	SuaraTotal          int `json:"suara_total" binding:"required"`
	PenggunaHakPilih    int `json:"pengguna_hak_pilih" binding:"required"`
	SuratSuaraDigunakan int `json:"surat_suara_digunakan" binding:"required"`

	JmlSuaraOcrList []int `json:"jml_suara_ocr_list" binding:"required"`
	JmlSuaraOmrList []int `json:"jml_suara_omr_list" binding:"required"`

	IdTps          int      `json:"id_tps" binding:"required"`
	JenisPemilihan int      `json:"jenis_pemilihan" binding:"required"`
	PayloadList    []string `json:"payload_list" binding:"required"`
	IdPaslonList   []int    `json:"id_paslon_list" binding:"required"`
	IdImageList    []int    `json:"IdImageList" binding:"required"`
	PdfUrl         string   `json:"pdf_url" binding:"required"`
}
