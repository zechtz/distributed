// Package registry provides registration for our services
package registry

// Registration struct
type Registration struct {
	ServiceName      ServiceName
	ServiceURL       string
	RequiredServices []ServiceName
	ServiceUpdateURL string
	HeartbeatURL     string
}

// ServiceName name of services that are being registered
type ServiceName string

const (
	// LogService the registration sercie
	LogService = ServiceName("LogService")
	// GradingService - the grading service
	GradingService = ServiceName("GradingService")
	// TeacherPortalService - the teacher portal service
	TeacherPortalService = ServiceName("TeacherPortalService")
)

type patchEntry struct {
	Name ServiceName
	URL  string
}

type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}
