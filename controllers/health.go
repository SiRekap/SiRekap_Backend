package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (h HealthController) GetHealthStatus(c *gin.Context) {
	c.String(http.StatusOK, "Connection success!")
}
