package controllers

import (
	"net/http"
	"sirekap/SiRekap_Backend/forms"
	"sirekap/SiRekap_Backend/models"

	"github.com/gin-gonic/gin"
)

var petugasTpsModel = new(models.PetugasTps)

type AuthController struct{}

func (a AuthController) Register(c *gin.Context) {
	var userRegisterData forms.UserRegisterData

	if err := c.BindJSON(&userRegisterData); err != nil {
		c.String(http.StatusBadRequest, "Data is not complete!")
	}

	petugasTps, err := petugasTpsModel.Register(userRegisterData)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, petugasTps)
	}
}
