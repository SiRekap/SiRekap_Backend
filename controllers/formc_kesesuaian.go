package controllers

import (
	"net/http"
	"sirekap/SiRekap_Backend/models"

	"github.com/gin-gonic/gin"
)

var formcKesesuaianModel = new(models.FormcKesesuaian)

type FormcKesesuaianController struct{}

func (f FormcKesesuaianController) SendFormCKesesuaian(c *gin.Context) {
	var formcKesesuaian models.FormcKesesuaian

	if err := c.BindJSON(&formcKesesuaian); err != nil {
		c.String(http.StatusBadRequest, "Data is not complete!")
		return
	}

	form, err := formcKesesuaianModel.SendFormCKesesuaian(formcKesesuaian)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, form)
	}
}
