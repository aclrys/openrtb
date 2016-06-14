package openrtb

// 5.16 Location Type
//
// The following table lists the options to indicate how the geographic information was determined.
type LocationType int8

const (
	LocationTypeGPSLocationServices LocationType = 1 // 1 GPS/Location Services
	LocationTypeIPAddress           LocationType = 2 // 2 IP Address
	LocationTypeUserProvided        LocationType = 3 // 3 User provided (e.g., registration data)
)
