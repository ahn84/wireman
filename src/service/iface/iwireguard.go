package iface

import (
	"context"
)

type WireguardSPI interface {
	CommonSPI
	AddPeers(ctx context.Context, peers []Peer) error
	RemovePeers(ctx context.Context, publicKeys []PublicKey) error
}

type CreateSessionReq struct {
	UserID    string
	PublicKey *string
}

type CreateSessionResp struct {
	SessionID *string `json:"session_id,omitempty"`
	UserID    *string `json:"user_id,omitempty"`
	PublicKey *string `json:"public_key,omitempty"`
	IPRange   *string `json:"ip_range,omitempty"`
}
