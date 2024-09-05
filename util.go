package useragent

import (
	"github.com/medama-io/go-useragent/internal"
)

// This adds a matching constant to a user agent struct.
func (ua *UserAgent) addMatch(result Result) bool {
	// Browsers
	if result.Type == internal.BrowserMatch && result.Precedence > ua.browserPrecedence {
		switch result.Match {
		case internal.Chrome:
			ua.browser = internal.Chrome
		case internal.Edge:
			ua.browser = internal.Edge
		case internal.Firefox:
			ua.browser = internal.Firefox
		case internal.IE:
			ua.browser = internal.IE
		case internal.Opera:
			ua.browser = internal.Opera
		case internal.OperaMini:
			ua.browser = internal.OperaMini
			ua.mobile = true
		case internal.Safari:
			ua.browser = internal.Safari
		case internal.Vivaldi:
			ua.browser = internal.Vivaldi
		case internal.Samsung:
			ua.browser = internal.Samsung
		case internal.Nintendo:
			ua.browser = internal.Nintendo
		case internal.YandexBrowser:
			ua.browser = internal.YandexBrowser
		}

		ua.browserPrecedence = result.Precedence
		return true
	}

	// Operating Systems
	if result.Type == internal.OSMatch && result.Precedence > ua.osPrecedence {
		switch result.Match {
		case internal.Android:
			ua.os = internal.Android
			// An older generic white-labeled variant of Chrome/Chromium on Android.
			if ua.browser == "" {
				ua.browser = internal.AndroidBrowser
				// Special case we set this as the precedence with this is zero
				// and can be overwritten by Safari.
				ua.browserPrecedence = internal.MatchPrecedenceMap[internal.Mobile]
			}
		case internal.ChromeOS:
			ua.os = internal.ChromeOS
			ua.desktop = true

		case internal.IOS:
			ua.os = internal.IOS
			if !ua.tablet {
				ua.mobile = true
			}
		case internal.Linux:
			ua.os = internal.Linux
			if !ua.tablet && !ua.tv {
				ua.desktop = true
			}
		case internal.MacOS:
			ua.os = internal.MacOS
			ua.desktop = true
		case internal.Windows:
			ua.os = internal.Windows
			ua.desktop = true
		}

		ua.osPrecedence = result.Precedence
		return true
	}

	// Types
	if result.Type == internal.TypeMatch && result.Precedence > ua.typePrecedence {
		switch result.Match {
		case internal.Desktop:
			ua.desktop = true
		case internal.Tablet:
			if ua.mobile {
				ua.mobile = false
			}
			ua.tablet = true
		case internal.Mobile, internal.MobileDevice:
			if !ua.tablet {
				ua.mobile = true
				ua.desktop = false
			}
		case internal.TV:
			ua.tv = true
		case internal.Bot:
			ua.bot = true
		}

		ua.typePrecedence = result.Precedence
		return true
	}

	return false
}
