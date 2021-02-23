// Package registry provides registration for our services
package registry

// Registration struct
type Registration struct {
	ServiceName ServiceName
	ServiceURL  string
}

// ServiceName name of services that are being registered
type ServiceName string

const (
	// LogService
	LogService = ServiceName("LogService")
)
