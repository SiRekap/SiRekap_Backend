package controllers

import (
	"net/http"
	"sirekap/SiRekap_Backend/forms"
	"sirekap/SiRekap_Backend/models"

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
