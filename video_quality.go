package openrtb

// 5.11 Video Quality
//
// The following table lists the options for the video quality. These values are defined by the IAB –
// http://www.iab.net/media/file/long-form-video-final.pdf
type VideoQuality int8

const (
	VideoQualityUnknown                VideoQuality = 0 // 0 Unknown
	VideoQualityProfessionallyProduced VideoQuality = 1 // 1 Professionally Produced
	VideoQualityProsumer               VideoQuality = 2 // 2 Prosumer
	VideoQualityUserGenerated          VideoQuality = 3 // 3 User Generated (UGC)
)
