package rest

import (
	"wgd/src/service/iface"

	"github.com/gin-gonic/gin"
)

type wgHandler struct {
	svc iface.WireguardSPI
}

// NewWireguardHandler creates a new instance of the Wireguard handler
func NewWireguardHandler(svc iface.WireguardSPI) *wgHandler {
	return &wgHandler{
		svc: svc,
	}
}

func (h *wgHandler) GetClients(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
