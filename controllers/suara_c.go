package controllers

import (
	"net/http"
	"sirekap/SiRekap_Backend/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var suaraCFinalModel = new(models.SuaraCFinal)
var suaraCProsesModel = new(models.SuaraCProses)

type SuaraCController struct{}

func (s SuaraCController) GetSuaraCFinal(c *gin.Context) {
	idVersi := c.Param("id_versi")
	if idVersi == "" {
		c.String(http.StatusBadRequest, "Id Versi is not provided!")
		return
	}

	integerIdVersi, err := strconv.Atoi(idVersi)
	if err != nil {
		c.String(http.StatusBadRequest, "Id Versi is not valid!")
		return
	}

	res, err := models.GetSuaraCFinal(integerIdVersi)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (s SuaraCController) SendSuaraCFinal(c *gin.Context) {
	var suaraCFinal models.SuaraCFinal

	if err := c.BindJSON(&suaraCFinal); err != nil {
		c.String(http.StatusBadRequest, "Data is not complete!")
		return
	}

	form, err := models.SendSuaraCFinal(suaraCFinal)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, form)
	}
}

func (s SuaraCController) GetSuaraCProses(c *gin.Context) {
	idImage := c.Param("id_image")
	if idImage == "" {
		c.String(http.StatusBadRequest, "Id Image is not provided!")
		return
	}

	integerIdImage, err := strconv.Atoi(idImage)
	if err != nil {
		c.String(http.StatusBadRequest, "Id Image is not valid!")
		return
	}

	res, err := models.GetSuaraCProses(integerIdImage)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (s SuaraCController) SendSuaraCProses(c *gin.Context) {
	var suaraCProses models.SuaraCProses

	if err := c.BindJSON(&suaraCProses); err != nil {
		c.String(http.StatusBadRequest, "Data is not complete!")
		return
	}

	form, err := models.SendSuaraCProses(suaraCProses)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, form)
	}
}
