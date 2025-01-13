package useragent

import (
	"strings"

	"github.com/medama-io/go-useragent/internal"
)

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
func (ua UserAgent) GetBrowser() string {
	return ua.browser
}

// GetOS returns the operating system name. If no OS is found, it returns an empty string.
func (ua UserAgent) GetOS() string {
	return ua.os
}

// GetVersion returns the browser version. If no version is found, it returns an empty string.
func (ua UserAgent) GetVersion() string {
	return string(ua.version[:ua.versionIndex])
}

// GetMajorVersion returns the major version of the browser. If no version is found, it returns an empty string.
func (ua UserAgent) GetMajorVersion() string {
	if ua.versionIndex == 0 {
		return ""
	}

	version := string(ua.version[:ua.versionIndex])

	return strings.Split(version, ".")[0]
}
