package openrtb

// 5.19 No-Bid Reason Codes
//
// The following table lists the options for a bidder to signal the exchange as to why it did not offer a bid
// for the impression.
type NoBidReasonCode int8

const (
	NoBidReasonCodeUnknownError             NoBidReasonCode = 0 // 0 Unknown Error
	NoBidReasonCodeTechnicalError           NoBidReasonCode = 1 // 1 Technical Error
	NoBidReasonCodeInvalidRequest           NoBidReasonCode = 2 // 2 Invalid Request
	NoBidReasonCodeKnownWebSpider           NoBidReasonCode = 3 // 3 Known Web Spider
	NoBidReasonCodeSuspectedNonHumanTraffic NoBidReasonCode = 4 // 4 Suspected Non-Human Traffic
	NoBidReasonCodeCloudDataCenterOrProxyIP NoBidReasonCode = 5 // 5 Cloud, Data center, or Proxy IP
	NoBidReasonCodeUnsupportedDevice        NoBidReasonCode = 6 // 6 Unsupported Device
	NoBidReasonCodeBlockedPublisherOrSite   NoBidReasonCode = 7 // 7 Blocked Publisher or Site
	NoBidReasonCodeUnmatchedUser            NoBidReasonCode = 8 // 8 Unmatched User
)
