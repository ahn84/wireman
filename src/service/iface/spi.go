package iface

import "context"

// CommonSPI defines a common Service Provider Interface (SPI) for typical application components
type CommonSPI interface {
	// Initialize sets up any necessary configurations and prepares the component for use
	Initialize(ctx context.Context) error
	// Start begins the operation of the component, making it ready to serve requests or perform its tasks
	Start(ctx context.Context) error
	// Stop gracefully stops the component, ensuring any in-progress operations are completed
	Stop(ctx context.Context) error
	// HealthCheck performs a health check on the component, returning an error if it is unhealthy
	HealthCheck(ctx context.Context) error
}
