package controllers

import (
	"net/http"
	"sirekap/SiRekap_Backend/forms"
	"sirekap/SiRekap_Backend/models"

	"github.com/gin-gonic/gin"
)

type FormcImageController struct{}

func (f FormcImageController) SendFormcImagePayload(c *gin.Context) {
	var formcImagePayload models.FormcImagePayload

	if err := c.BindJSON(&formcImagePayload); err != nil {
		c.String(http.StatusBadRequest, "Data is not complete!")
		return
	}

	form, err := formcImagePayload.SendFormcImagePayload()
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

	form, err := formcImage.SendFormcImage()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, form)
	}
}

func (f FormcImageController) SendFormcImageRaw(c *gin.Context) {
	var formcImageRaw forms.FormcImageRaw

	if err := c.BindJSON(&formcImageRaw); err != nil {
		c.String(http.StatusBadRequest, "Data is not complete!")
		return
	}

	form, err := models.SendFormcImageRaw(formcImageRaw)
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

	form, err := formcStatusData.SendFormcStatusData()
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

	form, err := formcStatusImage.SendFormcStatusImage()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, form)
	}
}
