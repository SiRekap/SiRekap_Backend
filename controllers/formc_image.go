package controllers

import (
	"net/http"
	"sirekap/SiRekap_Backend/models"

	"github.com/gin-gonic/gin"
)

var formcImagePayloadModel = new(models.FormcImagePayload)
var formcImageModel = new(models.FormcImage)
var formcStatusDataModel = new(models.FormcStatusData)
var formcStatusImageModel = new(models.FormcStatusImage)

type FormcImageController struct{}

func (f FormcImageController) SendFormcImagePayload(c *gin.Context) {
	var formcImagePayload models.FormcImagePayload

	if err := c.BindJSON(&formcImagePayload); err != nil {
		c.String(http.StatusBadRequest, "Data is not complete!")
		return
	}

	form, err := formcImagePayloadModel.SendFormcImagePayload(formcImagePayload)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, form)
	}
}

func (f FormcImageController) SendFormcImage(c *gin.Context) {
	var formcImage models.FormcImage

	if err := c.BindJSON(&formcImage); err != nil {
		c.String(http.StatusBadRequest, "Data is not complete!")
		return
	}

	form, err := formcImageModel.SendFormcImage(formcImage)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, form)
	}
}

func (f FormcImageController) SendFormcStatusData(c *gin.Context) {
	var formcStatusData models.FormcStatusData

	if err := c.BindJSON(&formcStatusData); err != nil {
		c.String(http.StatusBadRequest, "Data is not complete!")
		return
	}

	form, err := formcStatusDataModel.SendFormcStatusData(formcStatusData)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, form)
	}
}

func (f FormcImageController) SendFormcStatusImage(c *gin.Context) {
	var formcStatusImage models.FormcStatusImage

	if err := c.BindJSON(&formcStatusImage); err != nil {
		c.String(http.StatusBadRequest, "Data is not complete!")
		return
	}

	form, err := formcStatusImageModel.SendFormcStatusImage(formcStatusImage)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, form)
	}
}
