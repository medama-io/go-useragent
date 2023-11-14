package useragent

import (
	"sort"

	str "github.com/boyter/go-string"
)

const (
	// The following constants are used to determine agents.
	// Browsers
	Chrome    = "Chrome"
	Edge      = "Edge"
	Firefox   = "Firefox"
	IE        = "IE"
	Opera     = "Opera"
	OperaMini = "Mini"
	Safari    = "Safari"
	Vivaldi   = "Vivaldi"
	Samsung   = "SamsungBrowser"

	// Operating Systems
	Android  = "Android"
	ChromeOS = "ChromeOS"
	iOS      = "iOS"
	Linux    = "Linux"
	MacOS    = "MacOS"
	Windows  = "Windows NT"

	// Devices
	iPad   = "iPad"
	iPhone = "iPhone"
	iPod   = "iPod"

	// Types
	Desktop = "Desktop"
	Mobile  = "Mobile"
	Tablet  = "Tablet"
	TV      = "TV"
	Bot     = "Bot"
)

// MatchMap is a map of user agent types to their matching strings.
// These are the tokens saved into the trie when populating it.
var MatchMap = map[string][]string{
	// Browsers
	Chrome:    {Chrome},
	Edge:      {Edge, "Edg"},
	Firefox:   {Firefox},
	IE:        {"MSIE", "Trident"},
	Opera:     {Opera, "OPR"},
	OperaMini: {OperaMini},
	Safari:    {Safari},
	Vivaldi:   {Vivaldi},
	Samsung:   {Samsung},

	// Operating Systems
	Android:  {Android},
	ChromeOS: {"CrOS"},
	iOS:      {iPhone, iPad, iPod},
	Linux:    {Linux},
	MacOS:    {"Macintosh"},
	Windows:  {Windows},

	// Types
	Desktop: {Desktop, "Ubuntu", "Fedora"},
	Mobile:  {Mobile, "ONEPLUS", "Huawei", "HTC", "Galaxy", iPhone, iPod, "Windows Phone"},
	Tablet:  {Tablet, "Touch", iPad},
	TV:      {TV, "Large Screen", "Smart Display", "PLAYSTATION"},
	Bot:     {Bot, "bot", "Yahoo! Slurp", "LinkCheck", "QuickLook", "Haosou", "Yahoo Ad", "GoogleProber", "GoogleProducer", "Mediapartners", "Headless", "facebookexternalhit", "facebookcatalog"},
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
	Safari:    1, // Is always at the end of a Chrome user agent.
	Chrome:    2,
	Firefox:   3,
	IE:        4,
	Opera:     5,
	OperaMini: 6,
	Edge:      7,
	Vivaldi:   8,
	Samsung:   9,

	// Operating Systems
	Linux:    1,
	Android:  2,
	iOS:      3,
	ChromeOS: 4,
	MacOS:    5,
	Windows:  6,

	// Types
	Desktop: 1,
	Mobile:  2,
	Tablet:  3,
	TV:      4,
	Bot:     5,
}

type MatchResults struct {
	EndIndex int
	Match    string
	// 0: Unknown, 1: Browser, 2: OS, 3: Type
	MatchType  uint8
	Precedence uint8
}

// GetMatchType returns the match type of a match result using the MatchPrecedenceMap.
func GetMatchType(match string) uint8 {
	switch match {
	case Chrome, Edge, Firefox, IE, Opera, OperaMini, Safari, Vivaldi, Samsung:
		return 1
	case Android, ChromeOS, iOS, Linux, MacOS, Windows:
		return 2
	case Desktop, Mobile, Tablet, Bot:
		return 3
	default:
		return 0
	}
}

// MatchTokenIndexes finds the start and end indexes of necessary tokens
// that match a known browser, device, or OS. This is used to determine
// when to insert a result value into the trie.
func MatchTokenIndexes(ua string) []MatchResults {
	var results []MatchResults
	for key, match := range MatchMap {
		for _, m := range match {
			indexes := str.IndexAll(ua, m, -1)

			// Return the last match.
			if len(indexes) == 0 {
				continue
			}

			lastIndex := indexes[len(indexes)-1]

			// Check if key match doesn't already exist in results.
			// This is to prevent duplicate matches in the trie.
			exists := false
			for _, r := range results {
				if r.Match == key {
					exists = true
					break
				}
			}

			if !exists {
				matchType := GetMatchType(key)
				results = append(results, MatchResults{EndIndex: lastIndex[1], Match: key, MatchType: matchType, Precedence: MatchPrecedenceMap[key]})
			}
		}
	}

	// Sort the results by EndIndex in descending order.
	// This allows us to determine the first matching token in the user agent
	// when we iterate over it when populating the trie.
	sort.Slice(results, func(i, j int) bool {
		return results[i].EndIndex > results[j].EndIndex
	})

	return results
}

// This adds a matching constant to a user agent struct.
func (ua *UserAgent) addMatch(result *Result, existingPrecedence Precedence) {
	match := result.result
	precedence := result.precedence

	// Browsers
	if result.resultType == 1 && precedence > existingPrecedence.Browser {
		switch match {
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
		}

		ua.precedence.Browser = precedence
	}

	// Operating Systems
	if result.resultType == 2 && precedence > existingPrecedence.OS {
		switch match {
		case Android:
			ua.OS = Android
		case ChromeOS:
			ua.OS = ChromeOS
		case iOS:
			ua.OS = iOS
		case Linux:
			ua.OS = Linux
		case MacOS:
			ua.OS = MacOS
		case Windows:
			ua.OS = Windows
			ua.Desktop = true
		}

		ua.precedence.OS = precedence
	}

	// Types
	if result.resultType == 3 && precedence > existingPrecedence.Type {
		switch match {
		case Desktop, Windows, MacOS, Linux, ChromeOS:
			ua.Desktop = true
		case Tablet, iPad:
			if ua.Mobile {
				ua.Mobile = false
			}
			ua.Tablet = true
		case Mobile, iPhone, iPod, Android, OperaMini:
			if !ua.Tablet {
				ua.Mobile = true
			}
		case TV:
			ua.TV = true
		case Bot:
			ua.Bot = true
		}
	}
}
