package models

import (
	"errors"
	"sirekap/SiRekap_Backend/db"

	"gorm.io/gorm"
)

type (
	FormCAdministrasiHlmSatuProses struct {
		IdImage int `json:"id_image" binding:"required" gorm:"primaryKey"`

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
	}

	FormCAdministrasiHlmSatuFinal struct {
		IdVersi int  `json:"id_versi" binding:"required" gorm:"primaryKey"`
		IdImage int  `json:"id_image" binding:"required"`
		IsBenar bool `json:"is_benar" binding:"required"`
	}

	FormCAdministrasiHlmSatu struct {
		IdVersi int `json:"id_versi" binding:"required" gorm:"primaryKey"`
		IdImage int `json:"id_image" binding:"required"`

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

		IsBenar bool `json:"is_benar" binding:"required"`
	}

	FormCAdministrasiHlmDuaProses struct {
		IdImage int `json:"id_image" binding:"required" gorm:"primaryKey"`

		SuaraSah            int `json:"suara_sah" binding:"required"`
		SuaraTidakSah       int `json:"suara_tidak_sah" binding:"required"`
		SuaraTotal          int `json:"suara_total" binding:"required"`
		PenggunaHakPilih    int `json:"pengguna_hak_pilih" binding:"required"`
		SuratSuaraDigunakan int `json:"surat_suara_digunakan" binding:"required"`
	}

	FormCAdministrasiHlmDuaFinal struct {
		IdVersi int  `json:"id_versi" binding:"required" gorm:"primaryKey"`
		IdImage int  `json:"id_image" binding:"required"`
		IsBenar bool `json:"is_benar" binding:"required"`
	}

	FormCAdministrasiHlmDua struct {
		IdVersi int `json:"id_versi" binding:"required" gorm:"primaryKey"`
		IdImage int `json:"id_image" binding:"required"`

		SuaraSah            int `json:"suara_sah" binding:"required"`
		SuaraTidakSah       int `json:"suara_tidak_sah" binding:"required"`
		SuaraTotal          int `json:"suara_total" binding:"required"`
		PenggunaHakPilih    int `json:"pengguna_hak_pilih" binding:"required"`
		SuratSuaraDigunakan int `json:"surat_suara_digunakan" binding:"required"`

		IsBenar bool `json:"is_benar" binding:"required"`
	}
)

func (f FormCAdministrasiHlmSatuProses) GetFormCAdministrasiHlmSatuProses(idImage int) (FormCAdministrasiHlmSatuProses, error) {
	db := db.GetDB()

	form := FormCAdministrasiHlmSatuProses{
		IdImage: idImage,
	}

	result := db.First(&form)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return FormCAdministrasiHlmSatuProses{}, result.Error
	}

	return form, nil
}

func (f FormCAdministrasiHlmSatuProses) SendFormCAdministrasiHlmSatuProses(form FormCAdministrasiHlmSatuProses) (*FormCAdministrasiHlmSatuProses, error) {
	db := db.GetDB()

	db.Create(&form)

	return &form, nil
}

func (f FormCAdministrasiHlmSatuFinal) GetFormCAdministrasiHlmSatuFinal(idVersi int) (FormCAdministrasiHlmSatuFinal, error) {
	db := db.GetDB()

	form := FormCAdministrasiHlmSatuFinal{
		IdVersi: idVersi,
	}

	result := db.First(&form)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return FormCAdministrasiHlmSatuFinal{}, result.Error
	}

	return form, nil
}

func (f FormCAdministrasiHlmSatuFinal) SendFormCAdministrasiHlmSatuFinal(form FormCAdministrasiHlmSatuFinal) (*FormCAdministrasiHlmSatuFinal, error) {
	db := db.GetDB()

	db.Create(&form)

	return &form, nil
}

func (f FormCAdministrasiHlmSatu) GetFormCAdministrasiHlmSatu(idVersi int) (FormCAdministrasiHlmSatu, error) {
	db := db.GetDB()

	form := FormCAdministrasiHlmSatu{
		IdVersi: idVersi,
	}

	result := db.First(&form)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return FormCAdministrasiHlmSatu{}, result.Error
	}

	return form, nil
}

func (f FormCAdministrasiHlmSatu) SendFormCAdministrasiHlmSatu(form FormCAdministrasiHlmSatu) (*FormCAdministrasiHlmSatu, error) {
	db := db.GetDB()

	db.Create(&form)

	return &form, nil
}

func (f FormCAdministrasiHlmDuaProses) GetFormCAdministrasiHlmDuaProses(idImage int) (FormCAdministrasiHlmDuaProses, error) {
	db := db.GetDB()

	form := FormCAdministrasiHlmDuaProses{
		IdImage: idImage,
	}

	result := db.First(&form)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return FormCAdministrasiHlmDuaProses{}, result.Error
	}

	return form, nil
}

func (f FormCAdministrasiHlmDuaProses) SendFormCAdministrasiHlmDuaProses(form FormCAdministrasiHlmDuaProses) (*FormCAdministrasiHlmDuaProses, error) {
	db := db.GetDB()

	db.Create(&form)

	return &form, nil
}

func (f FormCAdministrasiHlmDuaFinal) GetFormCAdministrasiHlmDuaFinal(idVersi int) (FormCAdministrasiHlmDuaFinal, error) {
	db := db.GetDB()

	form := FormCAdministrasiHlmDuaFinal{
		IdVersi: idVersi,
	}

	result := db.First(&form)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return FormCAdministrasiHlmDuaFinal{}, result.Error
	}

	return form, nil
}

func (f FormCAdministrasiHlmDuaFinal) SendFormCAdministrasiHlmDuaFinal(form FormCAdministrasiHlmDuaFinal) (*FormCAdministrasiHlmDuaFinal, error) {
	db := db.GetDB()

	db.Create(&form)

	return &form, nil
}

func (f FormCAdministrasiHlmDua) GetFormCAdministrasiHlmDua(idVersi int) (FormCAdministrasiHlmDua, error) {
	db := db.GetDB()

	form := FormCAdministrasiHlmDua{
		IdVersi: idVersi,
	}

	result := db.First(&form)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return FormCAdministrasiHlmDua{}, result.Error
	}

	return form, nil
}

func (f FormCAdministrasiHlmDua) SendFormCAdministrasiHlmDua(form FormCAdministrasiHlmDua) (*FormCAdministrasiHlmDua, error) {
	db := db.GetDB()

	db.Create(&form)

	return &form, nil
}
