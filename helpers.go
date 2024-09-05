package useragent

import (
	"strings"
)

// Helper function to get the major version of the browser.
func (ua UserAgent) GetMajorVersion() string {
	if ua.version == "" {
		return ""
	}

	return strings.Split(ua.version, ".")[0]
}

func (ua UserAgent) IsDesktop() bool {
	return ua.desktop
}

func (ua UserAgent) IsMobile() bool {
	return ua.mobile
}

func (ua UserAgent) IsTablet() bool {
	return ua.tablet
}

func (ua UserAgent) IsTV() bool {
	return ua.tv
}

func (ua UserAgent) IsBot() bool {
	return ua.bot
}

func (ua UserAgent) GetBrowser() string {
	return ua.browser
}

func (ua UserAgent) GetOS() string {
	return ua.os
}

func (ua UserAgent) GetVersion() string {
	return ua.version
}
