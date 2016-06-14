package openrtb

// 5.10 Video Start Delay
//
// The following table lists the various options for the video start delay. If the start delay value is greater
// than 0, then the position is mid-roll and the value indicates the start delay.
type VideoStartDelay int64

const (
	VideoStartDelayMidRollMinInclusive VideoStartDelay = 1  // > 0 Mid-Roll (value indicates start delay in second)
	VideoStartDelayPreRoll             VideoStartDelay = 0  // 0 Pre-Roll
	VideoStartDelayGenericMidRoll      VideoStartDelay = -1 // -1 Generic Mid-Roll
	VideoStartDelayGenericPostRoll     VideoStartDelay = -2 // -2 Generic Post-Roll
)
