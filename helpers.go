package useragent

import (
	"strings"

	"github.com/medama-io/go-useragent/agents"
	"github.com/medama-io/go-useragent/internal"
)

// Browser returns the browser name. If no browser is found, it returns an empty string.
func (ua UserAgent) Browser() agents.Browser {
	return ua.browser.GetMatchBrowser()
}

// OS returns the operating system name. If no OS is found, it returns an empty string.
func (ua UserAgent) OS() agents.OS {
	return ua.os.GetMatchOS()
}

// Device returns the device type as a string.
func (ua UserAgent) Device() agents.Device {
	return ua.device.GetMatchDevice()
}

// BrowserVersion returns the browser version. If no version is found, it returns an empty string.
func (ua UserAgent) BrowserVersion() string {
	return string(ua.version[:ua.versionIndex])
}

// BrowserVersionMajor returns the major version of the browser. If no version is found, it returns an empty string.
func (ua UserAgent) BrowserVersionMajor() string {
	if ua.versionIndex == 0 {
		return ""
	}

	version := ua.BrowserVersion()

	return strings.Split(version, ".")[0]
}

// BrowserVersionMinor returns the minor version of the browser. If no version is found, it returns an empty string.
func (ua UserAgent) BrowserVersionMinor() string {
	if ua.versionIndex == 0 {
		return ""
	}

	version := ua.BrowserVersion()

	parts := strings.Split(version, ".")
	if len(parts) < 2 {
		return ""
	}

	return parts[1]
}

// BrowserVersionPatch returns the patch version of the browser. If no version is found, it returns an empty string.
func (ua UserAgent) BrowserVersionPatch() string {
	if ua.versionIndex == 0 {
		return ""
	}

	version := ua.BrowserVersion()

	parts := strings.Split(version, ".")
	if len(parts) < 3 {
		return ""
	}

	// Sometimes the patch version has a suffix, e.g. "1.2.3b".
	return strings.Join(parts[2:], ".")
}

// IsDesktop returns true if the user agent is a desktop browser.
func (ua UserAgent) IsDesktop() bool {
	return ua.device == internal.DeviceDesktop
}

// IsMobile returns true if the user agent is a mobile browser.
func (ua UserAgent) IsMobile() bool {
	return ua.device == internal.DeviceMobile
}

// IsTablet returns true if the user agent is a tablet browser.
func (ua UserAgent) IsTablet() bool {
	return ua.device == internal.DeviceTablet
}

// IsTV returns true if the user agent is a TV browser.
func (ua UserAgent) IsTV() bool {
	return ua.device == internal.DeviceTV
}

// IsBot returns true if the user agent is a bot.
func (ua UserAgent) IsBot() bool {
	return ua.device == internal.DeviceBot
}

// GetBrowser returns the browser name. If no browser is found, it returns an empty string.
//
// Deprecated: Use .Browser() instead.
func (ua UserAgent) GetBrowser() string {
	return string(ua.browser.GetMatchBrowser())
}

// GetOS returns the operating system name. If no OS is found, it returns an empty string.
//
// Deprecated: Use .OS() instead.
func (ua UserAgent) GetOS() string {
	return string(ua.os.GetMatchOS())
}

// GetDevice returns the device type as a string.
//
// Deprecated: Use .Device() instead.
func (ua UserAgent) GetDevice() string {
	return string(ua.device.GetMatchDevice())
}

// GetVersion returns the browser version. If no version is found, it returns an empty string.
//
// Deprecated: Use .BrowserVersion() instead.
func (ua UserAgent) GetVersion() string {
	return ua.BrowserVersion()
}

// GetMajorVersion returns the major version of the browser. If no version is found, it returns an empty string.
//
// Deprecated: Use .BrowserVersionMajor() instead.
func (ua UserAgent) GetMajorVersion() string {
	return ua.BrowserVersionMajor()
}
