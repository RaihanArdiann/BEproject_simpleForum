package memberships

import (
	"github.com/gin-gonic/gin"
)

type handler struct {
	*gin.Engine
}

func NewHandler(api *gin.Engine) *handler {
	return &handler{
		Engine: api,
	}
}

func (h *handler) RegisterRoute() {
	route := h.Group("membership")
	route.GET("/ping", h.ping)
}
