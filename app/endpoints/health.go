package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthEndpoint struct{}

func (h *HealthEndpoint) Status(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
