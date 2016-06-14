package openrtb

// 5.12 VAST Companion Types
//
// The following table lists the options to indicate markup types allowed for video companion ads. This
// table is derived from IAB VAST 2.0+. Refer to www.iab.net/vast/ for more information.
type VASTCompanionType int8

const (
	VASTCompanionTypeStaticResource VASTCompanionType = 1 // 1 Static Resource
	VASTCompanionTypeHTMLResource   VASTCompanionType = 2 // 2 HTML Resource
	VASTCompanionTypeIframeResource VASTCompanionType = 3 // 3 iframe Resource
)
