package orva

// AccessType for devices and users
type AccessType int32

const (
	// AnonAccess : New to the system, needs to be registerd
	AnonAccess AccessType = 0

	// NormalAccess : Basic user access (for regular profiles)
	NormalAccess AccessType = 1

	// AdminAccess : Admin access to prod systems
	AdminAccess AccessType = 2

	// DevAccess : Development access to non-prod sytstems
	DevAccess AccessType = 3

	// RootAccess : Access To All Systems
	RootAccess AccessType = 4
)
