package openrtb

// 5.14 Content Context
//
// The following table lists the various options for indicating the type of content in which the impression
// will appear.
//
// This OpenRTB table has values derived from the IAB Quality Assurance Guidelines (QAG). Practitioners
// should keep in sync with updates to the QAG values as published on IAB.net.
type ContentContext int8

const (
	ContentContextVideo       ContentContext = 1 // 1 Video (a video file or stream that is being watched by the user, including (Internet) television broadcasts)
	ContentContextGame        ContentContext = 2 // 2 Game (an interactive software game that is being played by the user)
	ContentContextMusic       ContentContext = 3 // 3 Music (an audio file or stream that is being listened to by the user, including (Internet) radio broadcasts)
	ContentContextApplication ContentContext = 4 // 4 Application (an interactive software application that is being used by the user)
	ContentContextText        ContentContext = 5 // 5 Text (a document that is primarily textual in nature that is being read or viewed by the user, including web page, eBook, or news article)
	ContentContextOther       ContentContext = 6 // 6 Other (content type unknown or the user is consuming content which does not fit into one of the categories above)
	ContentContextUnknown     ContentContext = 7 // 7 Unknown
)
