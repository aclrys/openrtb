package openrtb

// 5.13 Content Delivery Methods
//
// The following table lists the various options for the delivery of video content.
type ContentDeliveryMethod int8

const (
	ContentDeliveryMethodStreaming   ContentDeliveryMethod = 1 // 1 Streaming
	ContentDeliveryMethodProgressive ContentDeliveryMethod = 2 // 2 Progressive
)
