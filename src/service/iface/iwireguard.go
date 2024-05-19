package iface

import (
	"context"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

type Peer struct {
	wgtypes.Peer
}

type WireguardSPI interface {
	CommonSPI
	// AddPeer adds a new peer to the Wireguard interface
	AddPeer(ctx context.Context, peer *Peer) error
	// RemovePeer removes a peer from the Wireguard interface
	RemovePeer(ctx context.Context, peer *Peer) error
	// GetPeers returns a list of all peers currently connected to the Wireguard interface
	GetPeers(ctx context.Context) ([]*Peer, error)
	// GetPeer returns a specific peer connected to the Wireguard interface
	GetPeer(ctx context.Context, publicKey string) (*Peer, error)
}
