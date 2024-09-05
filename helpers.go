package useragent

import (
	"strings"
)

// IsDesktop returns true if the user agent is a desktop browser.
func (ua UserAgent) IsDesktop() bool {
	return ua.desktop
}

// IsMobile returns true if the user agent is a mobile browser.
func (ua UserAgent) IsMobile() bool {
	return ua.mobile
}

// IsTablet returns true if the user agent is a tablet browser.
func (ua UserAgent) IsTablet() bool {
	return ua.tablet
}

// IsTV returns true if the user agent is a TV browser.
func (ua UserAgent) IsTV() bool {
	return ua.tv
}

// IsBot returns true if the user agent is a bot.
func (ua UserAgent) IsBot() bool {
	return ua.bot
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
	return ua.version
}

// GetMajorVersion returns the major version of the browser. If no version is found, it returns an empty string.
func (ua UserAgent) GetMajorVersion() string {
	if ua.version == "" {
		return ""
	}

	return strings.Split(ua.version, ".")[0]
}
