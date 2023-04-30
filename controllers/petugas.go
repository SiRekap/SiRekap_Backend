package controllers

import (
	"net/http"
	"sirekap/SiRekap_Backend/forms"
	"sirekap/SiRekap_Backend/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var petugasTpsModel = new(models.PetugasTps)
var pemeriksaModel = new(models.Pemeriksa)

type PetugasController struct{}

func (p PetugasController) RegisterPetugas(c *gin.Context) {
	var petugasRegisterData forms.PetugasRegisterData

	if err := c.BindJSON(&petugasRegisterData); err != nil {
		c.String(http.StatusBadRequest, "Data is not complete!")
		return
	}

	petugasTps, err := petugasTpsModel.RegisterPetugas(petugasRegisterData)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, petugasTps)
	}
}

func (p PetugasController) RegisterPemeriksa(c *gin.Context) {
	var pemeriksaRegisterData forms.PemeriksaRegisterData

	if err := c.BindJSON(&pemeriksaRegisterData); err != nil {
		c.String(http.StatusBadRequest, "Data is not complete!")
		return
	}

	petugasTps, err := pemeriksaModel.RegisterPemeriksa(pemeriksaRegisterData)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, petugasTps)
	}
}

func (p PetugasController) GetPetugasTpsByIdPetugas(c *gin.Context) {

	idPetugas := c.Param("id_petugas")

	if idPetugas == "" {
		c.String(http.StatusBadRequest, "Id Petugas is not provided!")
		return
	}
	integerIdPetugas, err := strconv.Atoi(idPetugas)
	if err != nil {
		c.String(http.StatusBadRequest, "Id Petugas is not valid!")
		return
	}

	petugasTps, err := pemeriksaModel.GetPetugasTpsByIdPetugas(integerIdPetugas)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, petugasTps)
	}
}

func (p PetugasController) GetAllPemeriksaByTpsAndJenisPemilihan(c *gin.Context) {

	params := c.Request.URL.Query()

	idTps := params.Get("id_tps")

	if idTps == "" {
		c.String(http.StatusBadRequest, "Id TPS is not provided!")
		return
	}
	integerIdTps, err := strconv.Atoi(idTps)
	if err != nil {
		c.String(http.StatusBadRequest, "Id TPS is not valid!")
		return
	}

	jenisPemilihan := params.Get("jenis_pemilihan")

	if jenisPemilihan == "" {
		c.String(http.StatusBadRequest, "Jenis Pemilihan is not provided!")
		return
	}
	integerJenisPemilihan, err := strconv.Atoi(jenisPemilihan)
	if err != nil {
		c.String(http.StatusBadRequest, "Jenis Pemilihan is not valid!")
		return
	}

	petugasTpsList, err := pemeriksaModel.GetAllPemeriksaByTpsAndJenisPemilihan(integerIdTps, integerJenisPemilihan)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, petugasTpsList)
	}
}
