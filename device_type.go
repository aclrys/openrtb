package openrtb

// 5.17 Device Type
//
// The following table lists the type of device from which the impression originated.
//
// OpenRTB version 2.2 of the specification added distinct values for Mobile and Tablet. It is
// recommended that any bidder adding support for 2.2 treat a value of 1 as an acceptable alias of 4 & 5.
//
// This OpenRTB table has values derived from the IAB Quality Assurance Guidelines (QAG). Practitioners
// should keep in sync with updates to the QAG values as published on IAB.net.
type DeviceType int8

const (
	DeviceTypeMobileTablet     DeviceType = 1 // 1 Mobile/Tablet Version 2.0
	DeviceTypePersonalComputer DeviceType = 2 // 2 Personal Computer Version 2.0
	DeviceTypeConnectedTV      DeviceType = 3 // 3 Connected TV Version 2.0
	DeviceTypePhone            DeviceType = 4 // 4 Phone New for Version 2.2
	DeviceTypeTablet           DeviceType = 5 // 5 Tablet New for Version 2.2
	DeviceTypeConnectedDevice  DeviceType = 6 // 6 Connected Device New for Version 2.2
	DeviceTypeSetTopBox        DeviceType = 7 // 7 Set Top Box New for Version 2.2
)
