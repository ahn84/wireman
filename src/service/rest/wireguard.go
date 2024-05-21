package rest

import (
	"context"
	"net"
	"wgd/src/service/iface"

	"github.com/gin-gonic/gin"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
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

func (h *wgHandler) AddPeers(c *gin.Context) {
	_, cidr, _ := net.ParseCIDR("192.168.16.0/24")
	// SA/mNGO+Q9SoisQuTSxD40el2fpx490ZiA5k6CvhG1M=
	pub, _ := wgtypes.ParseKey("rgTQyhdwEH/KKNUFwqFBNx/ACfFJ3HUlygf3nPjhiAM=")
	h.svc.AddPeers(context.Background(), []iface.Peer{
		{AllowedIPs: []net.IPNet{*cidr}, PublicKey: iface.PublicKey{Key: pub}},
	})

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
