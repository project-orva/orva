package orva

// AccountType type of profile
type AccountType int32

const (
	// DeviceAccount account type that will be ran against the device accounts.
	DeviceAccount AccountType = 0

	// UserAccount account type that will be ran against the user accounts.
	UserAccount AccountType = 1
)