package openrtb

// 5.18 Connection Type
//
// The following table lists the various options for the type of device connectivity.
type ConnectionType int8

const (
	ConnectionTypeUnknown                          ConnectionType = 0 // 0 Unknown
	ConnectionTypeEthernet                         ConnectionType = 1 // 1 Ethernet
	ConnectionTypeWIFI                             ConnectionType = 2 // 2 WIFI
	ConnectionTypeCellularNetworkUnknownGeneration ConnectionType = 3 // 3 Cellular Network – Unknown Generation
	ConnectionTypeCellularNetwork2G                ConnectionType = 4 // 4 Cellular Network – 2G
	ConnectionTypeCellularNetwork3G                ConnectionType = 5 // 5 Cellular Network – 3G
	ConnectionTypeCellularNetwork4G                ConnectionType = 6 // 6 Cellular Network – 4G
)
