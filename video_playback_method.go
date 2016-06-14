package openrtb

// 5.9 Video Playback Methods
//
// The following table lists the various video playback methods.
type VideoPlaybackMethod int8

const (
	VideoPlaybackMethodAutoPlaySoundOn  VideoPlaybackMethod = 1 // 1 Auto-Play Sound On
	VideoPlaybackMethodAutoPlaySoundOff VideoPlaybackMethod = 2 // 2 Auto-Play Sound Off
	VideoPlaybackMethodClickToPlay      VideoPlaybackMethod = 3 // 3 Click-to-Play
	VideoPlaybackMethodMouseOver        VideoPlaybackMethod = 4 // 4 Mouse-Over
)
