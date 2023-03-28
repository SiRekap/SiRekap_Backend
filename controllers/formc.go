package controllers

import (
	"net/http"
	"sirekap/SiRekap_Backend/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var formCAdministrasiHlmSatuProsesModel = new(models.FormCAdministrasiHlmSatuProses)
var formCAdministrasiHlmSatuFinalModel = new(models.FormCAdministrasiHlmSatuFinal)
var formCAdministrasiHlmSatuModel = new(models.FormCAdministrasiHlmSatu)

var formCAdministrasiHlmDuaProsesModel = new(models.FormCAdministrasiHlmDuaProses)
var formCAdministrasiHlmDuaFinalModel = new(models.FormCAdministrasiHlmDuaFinal)
var formCAdministrasiHlmDuaModel = new(models.FormCAdministrasiHlmDua)

type FormCController struct{}

func (f FormCController) GetFormCAdministrasiHlmSatuProses(c *gin.Context) {
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

	res, err := formCAdministrasiHlmSatuProsesModel.GetFormCAdministrasiHlmSatuProses(integerIdImage)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (f FormCController) SendFormCAdministrasiHlmSatuProses(c *gin.Context) {
	var formCAdministrasiHlmSatuProses models.FormCAdministrasiHlmSatuProses

	if err := c.BindJSON(&formCAdministrasiHlmSatuProses); err != nil {
		c.String(http.StatusBadRequest, "Data is not complete!")
		return
	}

	res, err := formCAdministrasiHlmSatuProsesModel.SendFormCAdministrasiHlmSatuProses(formCAdministrasiHlmSatuProses)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (f FormCController) GetFormCAdministrasiHlmSatuFinal(c *gin.Context) {
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

	res, err := formCAdministrasiHlmSatuFinalModel.GetFormCAdministrasiHlmSatuFinal(integerIdVersi)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (f FormCController) SendFormCAdministrasiHlmSatuFinal(c *gin.Context) {
	var formCAdministrasiHlmSatuFinal models.FormCAdministrasiHlmSatuFinal

	if err := c.BindJSON(&formCAdministrasiHlmSatuFinal); err != nil {
		c.String(http.StatusBadRequest, "Data is not complete!")
		return
	}

	res, err := formCAdministrasiHlmSatuFinalModel.SendFormCAdministrasiHlmSatuFinal(formCAdministrasiHlmSatuFinal)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (f FormCController) GetFormCAdministrasiHlmSatu(c *gin.Context) {
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

	res, err := formCAdministrasiHlmSatuModel.GetFormCAdministrasiHlmSatu(integerIdVersi)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (f FormCController) SendFormCAdministrasiHlmSatu(c *gin.Context) {
	var formCAdministrasiHlmSatu models.FormCAdministrasiHlmSatu

	if err := c.BindJSON(&formCAdministrasiHlmSatu); err != nil {
		c.String(http.StatusBadRequest, "Data is not complete!")
		return
	}

	res, err := formCAdministrasiHlmSatuModel.SendFormCAdministrasiHlmSatu(formCAdministrasiHlmSatu)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (f FormCController) GetFormCAdministrasiHlmDuaProses(c *gin.Context) {
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

	res, err := formCAdministrasiHlmDuaProsesModel.GetFormCAdministrasiHlmDuaProses(integerIdImage)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (f FormCController) SendFormCAdministrasiHlmDuaProses(c *gin.Context) {
	var formCAdministrasiHlmDuaProses models.FormCAdministrasiHlmDuaProses

	if err := c.BindJSON(&formCAdministrasiHlmDuaProses); err != nil {
		c.String(http.StatusBadRequest, "Data is not complete!")
		return
	}

	res, err := formCAdministrasiHlmDuaProsesModel.SendFormCAdministrasiHlmDuaProses(formCAdministrasiHlmDuaProses)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (f FormCController) GetFormCAdministrasiHlmDuaFinal(c *gin.Context) {
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

	res, err := formCAdministrasiHlmDuaFinalModel.GetFormCAdministrasiHlmDuaFinal(integerIdVersi)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (f FormCController) SendFormCAdministrasiHlmDuaFinal(c *gin.Context) {
	var formCAdministrasiHlmDuaFinal models.FormCAdministrasiHlmDuaFinal

	if err := c.BindJSON(&formCAdministrasiHlmDuaFinal); err != nil {
		c.String(http.StatusBadRequest, "Data is not complete!")
		return
	}

	res, err := formCAdministrasiHlmDuaFinalModel.SendFormCAdministrasiHlmDuaFinal(formCAdministrasiHlmDuaFinal)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (f FormCController) GetFormCAdministrasiHlmDua(c *gin.Context) {
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

	res, err := formCAdministrasiHlmDuaModel.GetFormCAdministrasiHlmDua(integerIdVersi)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (f FormCController) SendFormCAdministrasiHlmDua(c *gin.Context) {
	var formCAdministrasiHlmDua models.FormCAdministrasiHlmDua

	if err := c.BindJSON(&formCAdministrasiHlmDua); err != nil {
		c.String(http.StatusBadRequest, "Data is not complete!")
		return
	}

	res, err := formCAdministrasiHlmDuaModel.SendFormCAdministrasiHlmDua(formCAdministrasiHlmDua)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}
