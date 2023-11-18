package useragent

import (
	"strings"
)

// Helper function to get the major version of the browser.
func (ua UserAgent) GetMajorVersion() string {
	if ua.Version == "" {
		return ""
	}

	return strings.Split(ua.Version, ".")[0]
}
