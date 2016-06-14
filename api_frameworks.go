package openrtb

// 5.6 API Frameworks
//
// The following table is a list of API frameworks supported by the publisher. Note that MRAID-1 is a
// subset of MRAID-2. In OpenRTB 2.1 and prior, value “3” was “MRAID”. However, not all MRAID capable
// APIs understand MRAID-2 features and as such the only safe interpretation of value “3” is MRAID-1. In
// OpenRTB 2.2, this was made explicit and MRAID-2 has been added as value “5”.
type APIFramework int8

const (
	APIFrameworkVPAID10 APIFramework = 1 // 1 VPAID 1.0
	APIFrameworkVPAID20 APIFramework = 2 // 2 VPAID 2.0
	APIFrameworkMRAID1  APIFramework = 3 // 3 MRAID-1
	APIFrameworkORMMA   APIFramework = 4 // 4 ORMMA
	APIFrameworkMRAID2  APIFramework = 5 // 5 MRAID-2
)
