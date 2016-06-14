package openrtb

// 5.4 Ad Position
//
// The following table specifies the position of the ad as a relative measure of visibility or prominence.
// This OpenRTB table has values derived from the IAB Quality Assurance Guidelines (QAG). Practitioners
// should keep in sync with updates to the QAG values as published on IAB.net. Values “3” – “6” apply to
// apps per the mobile addendum to QAG version 1.5.
type AdPosition int8

const (
	AdPositionUnknown                       AdPosition = 0 // 0 Unknown
	AdPositionAboveTheFold                  AdPosition = 1 // 1 Above the Fold
	AdPositionMayOrMayNotBeInitiallyVisible AdPosition = 2 // 2 DEPRECATED - May or may not be initially visible depending on screen size/resolution.
	AdPositionBelowTheFold                  AdPosition = 3 // 3 Below the Fold
	AdPositionHeader                        AdPosition = 4 // 4 Header
	AdPositionFooter                        AdPosition = 5 // 5 Footer
	AdPositionSidebar                       AdPosition = 6 // 6 Sidebar
	AdPositionFullScreen                    AdPosition = 7 // 7 Full Screen
)
