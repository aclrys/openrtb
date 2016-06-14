package openrtb

// 5.8 Video Bid Response Protocols
//
// The following table lists the options for video bid response protocols that could be supported by an
// exchange.
type VideoBidResponseProtocol int8

const (
	VideoBidResponseProtocolVAST10        VideoBidResponseProtocol = 1 // 1 VAST 1.0
	VideoBidResponseProtocolVAST20        VideoBidResponseProtocol = 2 // 2 VAST 2.0
	VideoBidResponseProtocolVAST30        VideoBidResponseProtocol = 3 // 3 VAST 3.0
	VideoBidResponseProtocolVAST10Wrapper VideoBidResponseProtocol = 4 // 4 VAST 1.0 Wrapper
	VideoBidResponseProtocolVAST20Wrapper VideoBidResponseProtocol = 5 // 5 VAST 2.0 Wrapper
	VideoBidResponseProtocolVAST30Wrapper VideoBidResponseProtocol = 6 // 6 VAST 3.0 Wrapper
)
