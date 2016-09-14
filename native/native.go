package native

// 4.1 Native Markup Request Object
//
// The Native Object defines the native advertising opportunity available for bid via this bid
// request. It must be included directly in the impression object if the impression offered for
// auction is a native ad format.
// The Default column dictates how optional parameters should be interpreted if explicit values
// are not provided.
type NativeMarkupRequest struct {

	// Field:
	//   ver
	// Scope:
	//   optional
	// Type:
	//   string
	// Default:
	//   1
	// Description:
	//   Version of the Native Markup
	//   version in use.
	Ver string `json:"ver,omitempty"`

	// Field:
	//   layout
	// Scope:
	//   recommended
	// Type:
	//   integer
	// Default:
	//   -
	// Description:
	//   The Layout ID of the native ad
	//   unit. See the table of Native
	//   Layout IDs below.
	Layout int8 `json:"layout,omitempty"`

	// Field:
	//   adunit
	// Scope:
	//   recommended
	// Type:
	//   integer
	// Default:
	//   -
	// Description:
	//   The Ad unit ID of the native ad
	//   unit. See the Table of Native Ad
	//   Unit IDs below for a list of
	//   supported core ad units.
	AdUnit int8 `json:"adunit,omitempty"`

	// Field:
	//   plcmtcnt
	// Scope:
	//   optional
	// Type:
	//   integer
	// Default:
	//   1
	// Description:
	//   The number of identical
	//   placements in this Layout. Refer
	//   to Section 8.1 Multi Placement
	//   Bid Requests.
	PlcmtCnt int8 `json:"plcmtcnt,omitempty"`

	// Field:
	//   seq
	// Scope:
	//   optional
	// Type:
	//   integer
	// Default:
	//   0
	// Description:
	//   xx (see the IAB Core Six layout
	//   types). 0 for the first ad, 1 for the
	//   second ad, and so on. This is not
	//   the sequence number of the
	//   content in the stream.
	Seq int8 `json:"seq,omitempty"`

	// Field:
	//   assets
	// Scope:
	//   required
	// Type:
	//   array of objects
	// Default:
	//   -
	// Description:
	//   An array of Asset Objects. Any
	//   bid must comply with the array
	//   of elements expressed by the
	//   Exchange.
	Assets []Asset `json:"assets"`

	// Field:
	//   ext
	// Scope:
	//   optional
	// Type:
	//   object
	// Default:
	//   -
	// Description:
	//   This object is a placeholder that
	//   may contain custom JSON agreed
	//   to by the parties to support
	//   flexibility beyond the standard
	//   defined in this specification.
	Ext Ext `json:"ext,omitempty"`
}

// 4.2 AssetObject
//
// The main container object for each asset requested or supported by Exchange on behalf of the
// rendering client. Any object that is required is to be flagged as such. Only one of the
// {title,,img,video,data} objects should be present in each object. All others should be
// null/absent. The id is to be unique within the AssetObject array so that the response can be
// aligned.
type Asset struct {

	// [1]: asset object may contain only one of title, img, data or video.

	// Field:
	//   id
	// Scope:
	//   required
	// Type:
	//   int
	// Default:
	//   -
	// Description:
	//   Unique asset ID, assigned by
	//   exchange. Typically a counter for
	//   the array.
	ID uint64 `json:"id"`

	// Field:
	//   required
	// Scope:
	//   optional
	// Type:
	//   int
	// Default:
	//   0
	// Description:
	//   Set to 1 if asset is required
	//   (exchange will not accept a bid
	//   without it).
	Required int8 `json:"required,omitempty"`

	// Field:
	//   title
	// Scope:
	//   optional [1]
	// Type:
	//   object
	// Default:
	//   -
	// Description:
	//   Title object for title assets. See
	//   Title Object definition.
	Title *Title `json:"title,omitempty"`

	// Field:
	//   img
	// Scope:
	//   optional [1]
	// Type:
	//   object
	// Default:
	//   -
	// Description:
	//   Image object for image assets.
	//   See Image Object definition.
	Img *Img `json:"img,omitempty"`

	// Field:
	//   video
	// Scope:
	//   optional [1]
	// Type:
	//   object
	// Default:
	//   -
	// Description:
	//   Video object for video assets.
	//   See the Video Object definition.
	//   Note that in-stream video ads
	//   are not part of Native. Native
	//   ads may contain a video as the
	//   ad creative itself.
	Video *Video `json:"video,omitempty"`

	// Field:
	//   data
	// Scope:
	//   optional [1]
	// Type:
	//   object
	// Default:
	//   -
	// Description:
	//   Data object for ratings, prices
	//   etc. See Data Object definition.
	Data *Data `json:"data,omitempty"`

	// Field:
	//   ext
	// Scope:
	//   optional
	// Type:
	//   object
	// Default:
	//   -
	// Description:
	//   This object is a placeholder that
	//   may contain custom JSON
	//   agreed to by the parties to
	//   support flexibility beyond the
	//   standard defined in this
	//   specification.
	Ext Ext `json:"ext,omitempty"`
}

// 4.3 Title Object
//
// The Title object is to be used for title element of the Native ad.
type Title struct {
	// Field:
	//   len
	// Scope:
	//   required
	// Type:
	//   integer
	// Default:
	//   -
	// Description:
	//   Maximum length of the text in
	//   the title element.
	Len uint64 `json:"len"`

	// Field:
	//   ext
	// Scope:
	//   optional
	// Type:
	//   object
	// Default:
	//   -
	// Description:
	//   This object is a placeholder that
	//   may contain custom JSON
	//   agreed to by the parties to
	//   support flexibility beyond the
	//   standard defined in this
	//   specification.
	Ext Ext `json:"ext,omitempty"`
}

// 4.4 Image Object
//
// The Image object to be used for all image elements of the Native ad such as Icons, Main Image,
// etc.
type Img struct {

	// Field:
	//   type
	// Scope:
	//   optional
	// Type:
	//   integer
	// Default:
	//   -
	// Description:
	//   Type ID of the image element
	//   supported by the publisher. The
	//   publisher can display this
	//   information in an appropriate
	//   format. See Table Image Asset
	//   Types for commonly used
	//   examples.
	Type int8 `json:"type,omitempty"`

	// Field:
	//   w
	// Scope:
	//   optional
	// Type:
	//   integer
	// Default:
	//   -
	// Description:
	//   Width of the image in pixels.
	W uint64 `json:"w,omitempty"`

	// Field:
	//   wmin
	// Scope:
	//   recommended
	// Type:
	//   integer
	// Default:
	//   -
	// Description:
	//   The minimum requested width
	//   of the image in pixels. This
	//   option should be used for any
	//   rescaling of images by the client.
	//   Either w or wmin should be
	//   transmitted. If only w is
	//   included, it should be considered
	//   an exact requirement.
	WMin uint64 `json:"w,omitempty"`

	// Field:
	//   h
	// Scope:
	//   optional
	// Type:
	//   integer
	// Default:
	//   -
	// Description:
	//   Height of the image in pixels.
	H uint64 `json:"h,omitempty"`

	// Field:
	//   hmin
	// Scope:
	//   recommended
	// Type:
	//   integer
	// Default:
	//   -
	// Description:
	//   The minimum requested height
	//   of the image in pixels. This
	//   option should be used for any
	//   rescaling of images by the client.
	//   Either h or hmin should be
	//   transmitted. If only h is included,
	//   it should be considered an exact
	//   requirement.
	HMin uint64 `json:"hmin,omitempty"`

	// Field:
	//   mimes
	// Scope:
	//   optional
	// Type:
	//   array of strings
	// Default:
	//   All
	//   types
	//   allowed
	// Description:
	//   Whitelist of content MIME types
	//   supported. Popular MIME types
	//   include, but are not limited to
	//   “image/jpg” “image/gif”.
	//   Each implementing Exchange
	//   should have their own list of
	//   supported types in the
	//   integration docs. See
	//   Wikipedia's MIME page for more
	//   information and links to all IETF
	//   RFCs.
	//   If blank, assume all types are
	//   allowed.
	Mimes []string `json:"mimes,omitempty"`

	// Field:
	//   ext
	// Scope:
	//   optional
	// Type:
	//   object
	// Default:
	//   -
	// Description:
	//   This object is a placeholder that
	//   may contain custom JSON
	//   agreed to by the parties to
	//   support flexibility beyond the
	//   standard defined in this
	//   specification.
	Ext Ext `json:"ext,omitempty"`
}

// 4.5 Video Object
//
// The video object to be used for all video elements supported in the Native Ad. This corresponds
// to the Video object of OpenRTB 2.3. Exchange implementers can impose their own specific
// restrictions. Here are the required attributes of the Video Object. For optional attributes please
// refer to OpenRTB 2.3.
type Video struct {

	// Field:
	//   mimes
	// Scope:
	//   required
	// Type:
	//   array of strings
	// Default:
	//   -
	// Description:
	//   Content MIME types supported.
	//   Popular MIME types include,but
	//   are not limited to “video/x-mswmv”
	//   for Windows Media, and
	//   “video/x-flv” for Flash Video.
	Mimes []string `json:"mimes"`

	// Field:
	//   minduration
	// Scope:
	//   required
	// Type:
	//   integer
	// Default:
	//   -
	// Description:
	//   Minimum video ad duration in
	//   seconds.
	MinDuration uint64 `json:"minduration"`

	// Field:
	//   maxduration
	// Scope:
	//   required
	// Type:
	//   integer
	// Default:
	//   -
	// Description:
	//   Maximum video ad duration in
	//   seconds.
	MaxDuration uint64 `json:"maxduration"`

	// Field:
	//   protocols
	// Scope:
	//   required
	// Type:
	//   array of integers
	// Default:
	//   -
	// Description:
	//   An array of video protocols the
	//   publisher can accept in the bid
	//   response. See OpenRTB 2.3
	//   Table 5.8 Video Bid Response
	//   Protocols for a list of possible
	//   values.
	Protocols []int8 `json:"protocols"`

	// Field:
	//   ext
	// Scope:
	//   optional
	// Type:
	//   object
	// Default:
	//   -
	// Description:
	//   This object is a placeholder that
	//   may contain custom JSON
	//   agreed to by the parties to
	//   support flexibility beyond the
	//   standard defined in this
	//   specification.
	Ext Ext `json:"ext,omitempty"`
}

// 4.6 Data Object
//
// The Data Object is to be used for all non-core elements of the native unit such as Ratings,
// Review Count, Stars, Download count, descriptions etc. It is also generic for future of Native
// elements not contemplated at the time of the writing of this document.
type Data struct {

	// Field:
	//   type
	// Scope:
	//   required
	// Type:
	//   integer
	// Default:
	//   -
	// Description:
	//   Type ID of the element
	//   supported by the publisher. The
	//   publisher can display this
	//   information in an appropriate
	//   format. See Table 7.3 Data
	//   Asset Types for commonly used
	//   examples.
	Type int8 `json:"type"`

	// Field:
	//   len
	// Scope:
	//   optional
	// Type:
	//   integer
	// Default:
	//   -
	// Description:
	//   Maximum length of the text in
	//   the element’s response.
	Len uint64 `json:"len,omitempty"`

	// Field:
	//   ext
	// Scope:
	//   optional
	// Type:
	//   object
	// Default:
	//   -
	// Description:
	//   This object is a placeholder that
	//   may contain custom JSON
	//   agreed to by the parties to
	//   support flexibility beyond the
	//   standard defined in this
	//   specification.
	Ext Ext `json:"ext,omitempty"`
}

// This object is a placeholder that
// may contain custom JSON agreed
// to by the parties to support
// flexibility beyond the standard
// defined in this specification.
type Ext interface{}
