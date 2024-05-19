package impl

import (
	"context"
	"wgd/src/service/iface"
)

type wireguardService struct{}

// NewWireguardService creates a new instance of the Wireguard service
func NewWireguardService() *wireguardService {
	return &wireguardService{}
}

// Initialize initializes the Wireguard service
func (s *wireguardService) Initialize(ctx context.Context) error {
	return nil
}

// Start starts the Wireguard service
func (s *wireguardService) Start(ctx context.Context) error {
	return nil
}

// Stop stops the Wireguard service
func (s *wireguardService) Stop(ctx context.Context) error {
	return nil
}

func (s *wireguardService) HealthCheck(ctx context.Context) error {
	return nil
}

// AddPeer adds a new peer to the Wireguard interface
func (s *wireguardService) AddPeer(ctx context.Context, p *iface.Peer) error {
	return nil
}

// RemovePeer removes a peer from the Wireguard interface
func (s *wireguardService) RemovePeer(ctx context.Context, p *iface.Peer) error {
	return nil
}

// GetPeers returns a list of all peers currently connected to the Wireguard interface
func (s *wireguardService) GetPeers(ctx context.Context) ([]*iface.Peer, error) {
	return nil, nil
}

// GetPeer returns a specific peer connected to the Wireguard interface
func (s *wireguardService) GetPeer(ctx context.Context, id string) (*iface.Peer, error) {
	return nil, nil
}
