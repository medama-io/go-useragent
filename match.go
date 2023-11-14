package useragent

import (
	"sort"

	str "github.com/boyter/go-string"
)

const (
	// The following constants are used to determine agents.
	// Browsers

	Chrome  = "Chrome"
	Edge    = "Edge"
	Firefox = "Firefox"
	IE      = "IE"
	Opera   = "Opera"
	Safari  = "Safari"
	Vivaldi = "Vivaldi"

	// Devices

	Android = "Android"
	iPad    = "iPad"
	iPhone  = "iPhone"
	iPod    = "iPod"

	// Operating Systems

	AndroidOS = "AndroidOS"
	ChromeOS  = "ChromeOS"
	iOS       = "iOS"
	Linux     = "Linux"
	MacOS     = "MacOS"
	Windows   = "Windows"

	// Other
	Desktop = "Desktop"
	Mobile  = "Mobile"
	Tablet  = "Tablet"
	Unknown = "Unknown"
)

// MatchMap is a map of user agent types to their matching strings.
var MatchMap = map[string][]string{
	// Browsers
	Chrome:  {"Chrome"},
	Edge:    {"Edge", "Edg"},
	Firefox: {"Firefox"},
	IE:      {"MSIE", "Trident"},
	Opera:   {"Opera"},
	Safari:  {"Safari"},
	Vivaldi: {"Vivaldi"},
	// Devices
	Android: {"Android"},
	iPad:    {"iPad"},
	iPhone:  {"iPhone"},
	iPod:    {"iPod"},
	// Operating Systems
	AndroidOS: {"Android"},
	ChromeOS:  {"CrOS"},
	iOS:       {"iPhone", "iPad", "iPod"},
	Linux:     {"Linux"},
	MacOS:     {"Macintosh"},
	Windows:   {"Windows"},
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
var MatchPrecedenceMap = map[string]int{
	// Browsers
	Safari:  1, // Is always at the end of a Chrome user agent.
	Chrome:  2,
	Firefox: 3,
	IE:      4,
	Opera:   5,
	Edge:    6,
	Vivaldi: 7,

	// Devices
	Android: 1,
	iPad:    1,
	iPhone:  1,
	iPod:    1,
	// Operating Systems
	AndroidOS: 1,
	ChromeOS:  1,
	iOS:       1,
	Linux:     1,
	MacOS:     1,
	Windows:   1,
}

type MatchResults struct {
	EndIndex   int
	Match      string
	Precedence int
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
				results = append(results, MatchResults{EndIndex: lastIndex[1], Match: key, Precedence: MatchPrecedenceMap[key]})
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
	if ua.Browser == "" || precedence > existingPrecedence.Browser {
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
		case Safari:
			ua.Browser = Safari
		}

		ua.precedence.Browser = precedence
	}

	// Devices
	if ua.Device == "" || precedence > existingPrecedence.Device {
		switch match {
		case Android:
			ua.Device = Android
		case iPad:
			ua.Device = iPad
		case iPhone:
			ua.Device = iPhone
		case iPod:
			ua.Device = iPod
		}

		ua.precedence.Device = precedence
	}

	// Operating Systems
	if ua.OS == "" || precedence > existingPrecedence.OS {
		switch match {
		case AndroidOS:
			ua.OS = AndroidOS
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
		}

		ua.precedence.OS = precedence
	}
}
