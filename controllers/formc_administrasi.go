package controllers

import (
	"net/http"
	"sirekap/SiRekap_Backend/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FormcAdministrasiController struct{}

func (f FormcAdministrasiController) GetFormcAdministrasiHlmSatuProses(c *gin.Context) {
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

	res, err := models.GetFormcAdministrasiHlmSatuProses(integerIdImage)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (f FormcAdministrasiController) SendFormcAdministrasiHlmSatuProses(c *gin.Context) {
	var formCAdministrasiHlmSatuProses models.FormcAdministrasiHlmSatuProses

	if err := c.BindJSON(&formCAdministrasiHlmSatuProses); err != nil {
		c.String(http.StatusBadRequest, "Data is not complete!")
		return
	}

	res, err := models.SendFormcAdministrasiHlmSatuProses(formCAdministrasiHlmSatuProses)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (f FormcAdministrasiController) GetFormcAdministrasiHlmSatuFinal(c *gin.Context) {
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

	res, err := models.GetFormcAdministrasiHlmSatuFinal(integerIdVersi)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (f FormcAdministrasiController) SendFormcAdministrasiHlmSatuFinal(c *gin.Context) {
	var formCAdministrasiHlmSatuFinal models.FormcAdministrasiHlmSatuFinal

	if err := c.BindJSON(&formCAdministrasiHlmSatuFinal); err != nil {
		c.String(http.StatusBadRequest, "Data is not complete!")
		return
	}

	res, err := models.SendFormcAdministrasiHlmSatuFinal(formCAdministrasiHlmSatuFinal)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// func (f FormcAdministrasiController) GetFormcAdministrasiHlmSatu(c *gin.Context) {
// 	idVersi := c.Param("id_versi")
// 	if idVersi == "" {
// 		c.String(http.StatusBadRequest, "Id Versi is not provided!")
// 		return
// 	}

// 	integerIdVersi, err := strconv.Atoi(idVersi)
// 	if err != nil {
// 		c.String(http.StatusBadRequest, "Id Versi is not valid!")
// 		return
// 	}

// 	res, err := formCAdministrasiHlmSatuModel.GetFormcAdministrasiHlmSatu(integerIdVersi)
// 	if err != nil {
// 		c.String(http.StatusBadRequest, err.Error())
// 		return
// 	} else {
// 		c.JSON(http.StatusOK, res)
// 	}
// }

// func (f FormcAdministrasiController) SendFormcAdministrasiHlmSatu(c *gin.Context) {
// 	var formCAdministrasiHlmSatu models.FormcAdministrasiHlmSatu

// 	if err := c.BindJSON(&formCAdministrasiHlmSatu); err != nil {
// 		c.String(http.StatusBadRequest, "Data is not complete!")
// 		return
// 	}

// 	res, err := formCAdministrasiHlmSatuModel.SendFormcAdministrasiHlmSatu(formCAdministrasiHlmSatu)
// 	if err != nil {
// 		c.String(http.StatusBadRequest, err.Error())
// 		return
// 	} else {
// 		c.JSON(http.StatusOK, res)
// 	}
// }

func (f FormcAdministrasiController) GetFormcAdministrasiHlmDuaProses(c *gin.Context) {
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

	res, err := models.GetFormcAdministrasiHlmDuaProses(integerIdImage)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (f FormcAdministrasiController) SendFormcAdministrasiHlmDuaProses(c *gin.Context) {
	var formCAdministrasiHlmDuaProses models.FormcAdministrasiHlmDuaProses

	if err := c.BindJSON(&formCAdministrasiHlmDuaProses); err != nil {
		c.String(http.StatusBadRequest, "Data is not complete!")
		return
	}

	res, err := models.SendFormcAdministrasiHlmDuaProses(formCAdministrasiHlmDuaProses)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (f FormcAdministrasiController) GetFormcAdministrasiHlmDuaFinal(c *gin.Context) {
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

	res, err := models.GetFormcAdministrasiHlmDuaFinal(integerIdVersi)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (f FormcAdministrasiController) SendFormcAdministrasiHlmDuaFinal(c *gin.Context) {
	var formCAdministrasiHlmDuaFinal models.FormcAdministrasiHlmDuaFinal

	if err := c.BindJSON(&formCAdministrasiHlmDuaFinal); err != nil {
		c.String(http.StatusBadRequest, "Data is not complete!")
		return
	}

	res, err := models.SendFormcAdministrasiHlmDuaFinal(formCAdministrasiHlmDuaFinal)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// func (f FormcAdministrasiController) GetFormcAdministrasiHlmDua(c *gin.Context) {
// 	idVersi := c.Param("id_versi")
// 	if idVersi == "" {
// 		c.String(http.StatusBadRequest, "Id Versi is not provided!")
// 		return
// 	}

// 	integerIdVersi, err := strconv.Atoi(idVersi)
// 	if err != nil {
// 		c.String(http.StatusBadRequest, "Id Versi is not valid!")
// 		return
// 	}

// 	res, err := formCAdministrasiHlmDuaModel.GetFormcAdministrasiHlmDua(integerIdVersi)
// 	if err != nil {
// 		c.String(http.StatusBadRequest, err.Error())
// 		return
// 	} else {
// 		c.JSON(http.StatusOK, res)
// 	}
// }

// func (f FormcAdministrasiController) SendFormcAdministrasiHlmDua(c *gin.Context) {
// 	var formCAdministrasiHlmDua models.FormcAdministrasiHlmDua

// 	if err := c.BindJSON(&formCAdministrasiHlmDua); err != nil {
// 		c.String(http.StatusBadRequest, "Data is not complete!")
// 		return
// 	}

// 	res, err := formCAdministrasiHlmDuaModel.SendFormcAdministrasiHlmDua(formCAdministrasiHlmDua)
// 	if err != nil {
// 		c.String(http.StatusBadRequest, err.Error())
// 		return
// 	} else {
// 		c.JSON(http.StatusOK, res)
// 	}
// }
