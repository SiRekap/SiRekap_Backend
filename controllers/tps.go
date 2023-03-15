package controllers

import (
	"net/http"
	"sirekap/SiRekap_Backend/models"

	"github.com/gin-gonic/gin"
)

var tpsModel = new(models.Tps)

type TpsController struct{}

func (t TpsController) GetTpsDetail(c *gin.Context) {
	params := c.Request.URL.Query()

	if !params.Has("id_tps") || params["id_tps"][0] == "" {
		c.String(http.StatusBadRequest, "Id TPS is not provided!")
	}

	tps, err := tpsModel.GetById(params["id_tps"][0])
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, tps)
	}
}
