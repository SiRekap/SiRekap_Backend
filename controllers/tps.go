package controllers

import (
	"net/http"
	"sirekap/SiRekap_Backend/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var tpsModel = new(models.Tps)

type TpsController struct{}

func (t TpsController) GetTpsDetail(c *gin.Context) {
	idTps := c.Param("id_tps")

	if idTps == "" {
		c.String(http.StatusBadRequest, "Id TPS is not provided!")
		return
	}
	integerIdTps, err := strconv.Atoi(idTps)
	if err != nil {
		c.String(http.StatusBadRequest, "Id TPS is not valid!")
		return
	}

	tps, err := tpsModel.GetById(integerIdTps)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, tps)
	}
}
