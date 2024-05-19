package configuration

import (
	"wgd/src/service/iface"
)

// AppContext will hold reference to service interfaces that will eventually be injected into this layer on initialization
type AppContext struct {
	HttpSvc   iface.CommonSPI
	Wireguard iface.WireguardSPI
}
