package iface

import (
	"net"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

type PrivateKey struct {
	wgtypes.Key
}

type PublicKey = PrivateKey

type Peer struct {
	PublicKey  PublicKey
	AllowedIPs []net.IPNet
}

// GeneratePrivateKey generates a new Wireguard private key
func GeneratePrivateKey() (PrivateKey, error) {
	privateKey, err := wgtypes.GeneratePrivateKey()
	return PrivateKey{privateKey}, err
}

func ValidatePublicKey(k string) error {
	// validate wireguard public key
	_, err := wgtypes.ParseKey(k)
	return err
}
