package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-clean-architecture/internal/response"
)

type HealthCheckController struct {}

func NewHealthCheckController() *HealthCheckController {
	return &HealthCheckController{}
}

//	@Summary	Health Check
//	@Success	200	{object}	response.HealthCheckResponse{}	"OK"
//	@Router		/healthz [get]
func (_ *HealthCheckController) HealthCheck(c *gin.Context) {
	response := response.HealthCheckResponse{Message: "OK"}

	c.JSON(http.StatusOK, response)
}