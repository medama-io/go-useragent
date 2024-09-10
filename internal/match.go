package internal

import (
	"sort"

	str "github.com/boyter/go-string"
)

const (
	// These are enum constants for the match type.
	BrowserMatch = 1
	OSMatch      = 2
	TypeMatch    = 3
	VersionMatch = 4
	UnknownMatch = 0

	// The following constants are used to determine agents.

	// Browsers.
	// There is no match token for this, but the absence of any browser token paired with Android is a good indicator of this browser.
	AndroidBrowser = "Android Browser"
	Chrome         = "Chrome"
	Edge           = "Edge"
	Firefox        = "Firefox"
	IE             = "IE"
	Opera          = "Opera"
	OperaMini      = "Mini"
	Safari         = "Safari"
	Vivaldi        = "Vivaldi"
	Samsung        = "Samsung Browser"
	Falkon         = "Falkon"
	Nintendo       = "Nintendo Browser"
	YandexBrowser  = "Yandex Browser"

	// Operating Systems.
	Android  = "Android"
	ChromeOS = "ChromeOS"
	IOS      = "iOS"
	Linux    = "Linux"
	OpenBSD  = "OpenBSD"
	MacOS    = "MacOS"
	Windows  = "Windows"

	// Types.
	Desktop = "Desktop"
	Mobile  = "Mobile"
	// We need a separate type for mobile devices since some user agents use "Mobile/"
	// appended with a device ID. We need to handle these separately to strip those IDs
	// out.
	MobileDevice = "MobileDevice"
	Tablet       = "Tablet"
	TV           = "TV"
	Bot          = "Bot"

	// Version.
	Version = "Version"
)

// matchMap is a map of user agent types to their matching strings.
// These are the tokens saved into the trie when populating it.
//
// Matching tokens are ordered by precedence. The first match is the
// most important match.
var matchMap = map[string][]string{
	// Browsers
	Chrome:        {"CriOS", Chrome},
	Edge:          {"EdgiOS", Edge, "Edg"},
	Firefox:       {"FxiOS", Firefox},
	IE:            {"MSIE", "Trident"},
	Opera:         {"OPiOS", "OPR", Opera},
	OperaMini:     {OperaMini},
	Safari:        {Safari, "AppleWebKit"},
	Vivaldi:       {Vivaldi},
	Samsung:       {"SamsungBrowser"},
	Falkon:        {Falkon},
	Nintendo:      {"NintendoBrowser"},
	YandexBrowser: {"YaBrowser"},

	// Operating Systems
	Android:  {Android},
	ChromeOS: {"CrOS"},
	IOS:      {"iPhone", "iPad", "iPod"},
	Linux:    {Linux, "Ubuntu", "Fedora"},
	OpenBSD:  {OpenBSD},
	MacOS:    {"Macintosh"},
	Windows:  {"Windows NT", "WindowsNT"},

	// Types
	Desktop:      {Desktop, "Ubuntu", "Fedora"},
	Mobile:       {Mobile},
	MobileDevice: {"ONEPLUS", "Huawei", "HTC", "Galaxy", "iPhone", "iPod", "Windows Phone", "WindowsPhone", "LG"},
	Tablet:       {Tablet, "Touch", "iPad", "Nintendo Switch", "NintendoSwitch", "Kindle"},
	TV:           {TV, "Large Screen", "LargeScreen", "Smart Display", "SmartDisplay", "PLAYSTATION", "PlayStation", "ADT-2", "ADT-1", "CrKey", "Roku", "AFT", "Web0S", "Nexus Player", "Xbox", "XBOX", "Nintendo WiiU", "NintendoWiiU"},
	Bot:          {Bot, "HeadlessChrome", "bot", "Slurp", "LinkCheck", "QuickLook", "Haosou", "Yahoo Ad", "YahooAd", "Google", "Mediapartners", "Headless", "facebookexternalhit", "facebookcatalog", "Baidu"},

	// Version
	Version: {Version},
}

// MatchPrecedenceMap is a map of user agent types to their importance
// in determining what is the actual browser/device/OS being used.
//
// For example, Chrome user agents also contain the string "Safari" at
// the end of the user agent. This means that if we only check for
// "Safari" first, we will incorrectly determine the browser to be Safari
// instead of Chrome.
//
// By setting a precedence, we can determine which match is more important
// and use that as the final result.
var MatchPrecedenceMap = map[string]uint8{
	// Browsers
	Safari:         1, // Is always at the end of a Chrome user agent.
	AndroidBrowser: 2,
	Chrome:         3,
	Firefox:        4,
	IE:             5,
	Opera:          6,
	OperaMini:      7,
	Edge:           8,
	Vivaldi:        9,
	Samsung:        10,
	Falkon:         11,
	Nintendo:       12,
	YandexBrowser:  13,

	// Operating Systems
	Linux:    1,
	Android:  2,
	IOS:      3,
	OpenBSD:  4,
	ChromeOS: 5,
	MacOS:    6,
	Windows:  7,

	// Types
	Desktop:      1,
	Mobile:       2,
	MobileDevice: 3,
	Tablet:       4,
	TV:           5,
	Bot:          6,
}

// MatchResults contains the information from MatchTokenIndexes.
type MatchResults struct {
	EndIndex int
	Match    string
	// 0: Unknown, 1: Browser, 2: OS, 3: Type
	MatchType uint8
	// Precedence value for each result type to determine which result should be overwritten.
	// Higher values overwrite lower values.
	Precedence uint8
}

// GetMatchType returns the match type of a match result using the MatchPrecedenceMap.
func GetMatchType(match string) uint8 {
	switch match {
	case Chrome, Edge, Firefox, IE, Opera, OperaMini, Safari, Vivaldi, Samsung, Falkon, Nintendo, YandexBrowser:
		return BrowserMatch
	case Android, ChromeOS, IOS, Linux, OpenBSD, MacOS, Windows:
		return OSMatch
	case Desktop, Mobile, MobileDevice, Tablet, Bot, TV:
		return TypeMatch
	case Version:
		return VersionMatch
	default:
		return UnknownMatch
	}
}

// MatchTokenIndexes finds the start and end indexes of necessary tokens
// that match a known browser, device, or OS. This is used to determine
// when to insert a result value into the trie.
func MatchTokenIndexes(ua string) []MatchResults {
	var results []MatchResults
	exists := make(map[string]bool)
	for key, match := range matchMap {
		for _, m := range match {
			// Check if key match doesn't already exist in results.
			// This is to prevent duplicate matches in the trie.
			if exists[key] {
				continue
			}

			indexes := str.IndexAll(ua, m, -1)

			// Return the last match.
			if len(indexes) == 0 {
				continue
			}

			lastIndex := indexes[len(indexes)-1]

			// Add the match to the results.
			matchType := GetMatchType(key)
			results = append(results, MatchResults{EndIndex: lastIndex[1], Match: key, MatchType: matchType, Precedence: MatchPrecedenceMap[key]})
			exists[key] = true
		}
	}

	// Sort the results by EndIndex in descending order.
	// This allows us to determine the first matching token in the user agent
	// when we iterate over it when populating the trie.
	//
	// Some tokens may have the same EndIndex, so we need to sort by Match key
	// to make it deterministic.
	sort.Slice(results, func(i, j int) bool {
		if results[i].EndIndex == results[j].EndIndex {
			return results[i].Match < results[j].Match
		}
		return results[i].EndIndex > results[j].EndIndex
	})

	return results
}
