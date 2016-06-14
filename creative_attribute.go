package openrtb

// 5.3 Creative Attributes
//
// The following table specifies a standard list of creative attributes that can describe an ad being served or
// serve as restrictions of thereof.
type CreativeAttribute int8

const (
	CreativeAttributeAudioAdAutoPlay                                CreativeAttribute = 1  // 1 Audio Ad (Auto-Play)
	CreativeAttributeAudioAdUserInitiated                           CreativeAttribute = 2  // 2 Audio Ad (User Initiated)
	CreativeAttributeExpandableAutomatic                            CreativeAttribute = 3  // 3 Expandable (Automatic)
	CreativeAttributeExpandableUserInitiatedClick                   CreativeAttribute = 4  // 4 Expandable (User Initiated - Click)
	CreativeAttributeExpandableUserInitiatedRollover                CreativeAttribute = 5  // 5 Expandable (User Initiated - Rollover)
	CreativeAttributeInBannerVideoAdAutoPlay                        CreativeAttribute = 6  // 6 In-Banner Video Ad (Auto-Play)
	CreativeAttributeInBannerVideoAdUserInitiated                   CreativeAttribute = 7  // 7 In-Banner Video Ad (User Initiated)
	CreativeAttributePop                                            CreativeAttribute = 8  // 8 Pop (e.g., Over, Under, or Upon Exit)
	CreativeAttributeProvocativeOrSuggestiveImagery                 CreativeAttribute = 9  // 9 Provocative or Suggestive Imagery
	CreativeAttributeShakyFlashingFlickeringExtremeAnimationSmileys CreativeAttribute = 10 // 10 Shaky, Flashing, Flickering, Extreme Animation, Smileys
	CreativeAttributeSurveys                                        CreativeAttribute = 11 // 11 Surveys
	CreativeAttributeTextOnly                                       CreativeAttribute = 12 // 12 Text Only
	CreativeAttributeUserInteractive                                CreativeAttribute = 13 // 13 User Interactive (e.g., Embedded Games)
	CreativeAttributeWindowsDialogOrAlertStyle                      CreativeAttribute = 14 // 14 Windows Dialog or Alert Style
	CreativeAttributeHasAudioOnOffButton                            CreativeAttribute = 15 // 15 Has Audio On/Off Button
	CreativeAttributeAdCanBeSkipped                                 CreativeAttribute = 16 // 16 Ad Can be Skipped (e.g., Skip Button on Pre-Roll Video)
)
