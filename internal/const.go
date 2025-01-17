package internal

type (
	// Match is an enum for the match type.
	Match uint8
	// Device is an enum for the device type.
	Device uint8
)

const (
	MatchUnknown Match = iota
	MatchBrowser
	MatchOS
	MatchType
	MatchVersion

	DeviceUnknown Device = iota
	DeviceDesktop
	DeviceMobile
	DeviceTablet
	DeviceTV
	DeviceBot

	DeviceUnknownStr = "Unknown"
	DeviceDesktopStr = "Desktop" 
	DeviceMobileStr = "Mobile"
	DeviceTabletStr = "Tablet"
	DeviceTVStr = "TV"
	DeviceBotStr = "Bot"

)



