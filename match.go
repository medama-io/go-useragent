package useragent

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
	Chrome    = "Chrome"
	Edge      = "Edge"
	Firefox   = "Firefox"
	IE        = "IE"
	Opera     = "Opera"
	OperaMini = "Mini"
	Safari    = "Safari"
	Vivaldi   = "Vivaldi"
	Samsung   = "SamsungBrowser"
	Nintendo  = "NintendoBrowser"

	// Operating Systems.
	Android  = "Android"
	ChromeOS = "ChromeOS"
	IOS      = "iOS"
	Linux    = "Linux"
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
	Chrome:    {"CriOS", Chrome},
	Edge:      {"EdgiOS", Edge, "Edg"},
	Firefox:   {"FxiOS", Firefox},
	IE:        {"MSIE", "Trident"},
	Opera:     {"OPiOS", "OPR", Opera},
	OperaMini: {OperaMini},
	Safari:    {Safari},
	Vivaldi:   {Vivaldi},
	Samsung:   {Samsung},
	Nintendo:  {Nintendo},

	// Operating Systems
	Android:  {Android},
	ChromeOS: {"CrOS"},
	IOS:      {"iPhone", "iPad", "iPod"},
	Linux:    {Linux},
	MacOS:    {"Macintosh"},
	Windows:  {"Windows NT", "WindowsNT"},

	// Types
	Desktop:      {Desktop, "Ubuntu", "Fedora"},
	Mobile:       {Mobile},
	MobileDevice: {"ONEPLUS", "Huawei", "HTC", "Galaxy", "iPhone", "iPod", "Windows Phone", "WindowsPhone", "LG"},
	Tablet:       {Tablet, "Touch", "iPad", "Nintendo Switch", "NintendoSwitch", "Kindle"},
	TV:           {TV, "Large Screen", "LargeScreen", "Smart Display", "SmartDisplay", "PLAYSTATION", "PlayStation", "ADT-2", "ADT-1", "CrKey", "Roku", "AFT", "Web0S", "Nexus Player", "Xbox", "XBOX", "Nintendo WiiU", "NintendoWiiU"},
	Bot:          {Bot, "bot", "Slurp", "LinkCheck", "QuickLook", "Haosou", "Yahoo Ad", "YahooAd", "GoogleProber", "GoogleProducer", "Mediapartners", "Headless", "facebookexternalhit", "facebookcatalog"},

	// Version
	Version: {Version},
}

// matchPrecedenceMap is a map of user agent types to their importance
// in determining what is the actual browser/device/OS being used.
//
// For example, Chrome user agents also contain the string "Safari" at
// the end of the user agent. This means that if we only check for
// "Safari" first, we will incorrectly determine the browser to be Safari
// instead of Chrome.
//
// By setting a precedence, we can determine which match is more important
// and use that as the final result.
var matchPrecedenceMap = map[string]uint8{
	// Browsers
	Safari:    1, // Is always at the end of a Chrome user agent.
	Chrome:    2,
	Firefox:   3,
	IE:        4,
	Opera:     5,
	OperaMini: 6,
	Edge:      7,
	Vivaldi:   8,
	Samsung:   9,
	Nintendo:  10,

	// Operating Systems
	Linux:    1,
	Android:  2,
	IOS:      3,
	ChromeOS: 4,
	MacOS:    5,
	Windows:  6,

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
	// Higher values are overwrite lower values.
	Precedence uint8
}

// GetMatchType returns the match type of a match result using the MatchPrecedenceMap.
func GetMatchType(match string) uint8 {
	switch match {
	case Chrome, Edge, Firefox, IE, Opera, OperaMini, Safari, Vivaldi, Samsung, Nintendo:
		return BrowserMatch
	case Android, ChromeOS, IOS, Linux, MacOS, Windows:
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
			results = append(results, MatchResults{EndIndex: lastIndex[1], Match: key, MatchType: matchType, Precedence: matchPrecedenceMap[key]})
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

// This adds a matching constant to a user agent struct.
func (ua *UserAgent) addMatch(result Result) bool {
	// Browsers
	if result.Type == BrowserMatch && result.Precedence > ua.precedence.Browser {
		switch result.Match {
		case Chrome:
			ua.Browser = Chrome
		case Edge:
			ua.Browser = Edge
		case Firefox:
			ua.Browser = Firefox
		case IE:
			ua.Browser = IE
		case Opera:
			ua.Browser = Opera
		case OperaMini:
			ua.Browser = OperaMini
			ua.Mobile = true
		case Safari:
			ua.Browser = Safari
		case Vivaldi:
			ua.Browser = Vivaldi
		case Samsung:
			ua.Browser = Samsung
		case Nintendo:
			ua.Browser = Nintendo
		}

		ua.precedence.Browser = result.Precedence
		return true
	}

	// Operating Systems
	if result.Type == OSMatch && result.Precedence > ua.precedence.OS {
		switch result.Match {
		case Android:
			ua.OS = Android
		case ChromeOS:
			ua.OS = ChromeOS
			ua.Desktop = true
		case IOS:
			ua.OS = IOS
			if !ua.Tablet {
				ua.Mobile = true
			}
		case Linux:
			ua.OS = Linux
			if !ua.Tablet && !ua.TV {
				ua.Desktop = true
			}
		case MacOS:
			ua.OS = MacOS
			ua.Desktop = true
		case Windows:
			ua.OS = Windows
			ua.Desktop = true
		}

		ua.precedence.OS = result.Precedence
		return true
	}

	// Types
	if result.Type == TypeMatch && result.Precedence > ua.precedence.Type {
		switch result.Match {
		case Desktop:
			ua.Desktop = true
		case Tablet:
			if ua.Mobile {
				ua.Mobile = false
			}
			ua.Tablet = true
		case Mobile, MobileDevice:
			if !ua.Tablet {
				ua.Mobile = true
			}
		case TV:
			ua.TV = true
		case Bot:
			ua.Bot = true
		}

		ua.precedence.Type = result.Precedence
		return true
	}

	return false
}
