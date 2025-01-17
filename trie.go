package useragent

import (
	"slices"

	"github.com/medama-io/go-useragent/internal"
)

const (
	// This is the number of rune trie children to store in the array
	// before switching to a map. Smaller arrays are faster to iterate
	// and use less memory.
	//
	// This is an arbitrary number, but seemed to perform well in benchmarks.
	maxChildArraySize = 64
)

type resultItem struct {
	Match string
	// 0: Unknown, 1: Browser, 2: OS, 3: Type
	Type internal.Match
	// Precedence value for each result type to determine which result
	// should be overwritten.
	Precedence uint8
}

type childNode struct {
	node *RuneTrie
	r    rune
}

// RuneTrie is a trie of runes with string keys and interface{} values.
type RuneTrie struct {
	// childrenArr is an array of pointers to the next RuneTrie node. This
	// is used when the number of children is small to improve performance
	// and reduce memory usage.
	childrenArr []childNode
	// childrenMap is a map of runes to the next RuneTrie node. This is used
	// when the number of children exceeds the diminishing returns of the
	// childrenArr array.
	childrenMap map[rune]*RuneTrie
	result      []resultItem
}

// NewRuneTrie allocates and returns a new *RuneTrie.
func NewRuneTrie() *RuneTrie {
	return new(RuneTrie)
}

// Get returns the value stored at the given key. Returns nil for internal
// nodes or for nodes with a value of nil.
func (trie *RuneTrie) Get(key string) UserAgent {
	node := trie
	var ua UserAgent

	// Flag to indicate if we are currently iterating over a version number.
	var isVersion bool
	// Number of runes to skip when iterating over the trie. This is used
	// to skip over version numbers or language codes.
	var skipCount uint8
	// Skip until we encounter whitespace.
	var skipUntilWhitespace bool
	// Skip until we encounter a closing parenthesis, used for skipping over device IDs.
	var skipUntilClosingParenthesis bool

	for i, r := range key {
		if skipUntilWhitespace {
			if r == ' ' {
				skipUntilWhitespace = false
			} else {
				continue
			}
		}

		if skipCount > 0 {
			skipCount--
			continue
		}

		if skipUntilClosingParenthesis {
			if r == ')' {
				skipUntilClosingParenthesis = false
			} else {
				continue
			}
		}

		if isVersion {
			// If we encounter any unknown characters, we can assume the version number is over.
			if !internal.IsDigit(r) && r != '.' {
				isVersion = false
			} else {
				// Add to rune buffer.
				if ua.versionIndex < cap(ua.version) {
					ua.version[ua.versionIndex] = r
					ua.versionIndex++
				}
				continue
			}
		}

		// Strip any other version numbers from other products to get more hits to the trie.
		//
		// Also do not use a switch here as Go does not generate a jump table for switch
		// statements with no integral constants. Benchmarking shows that ops go down
		// if we try to migrate statements like this to a switch.
		if internal.IsDigit(r) || (r == '.' && len(key) > i+1 && internal.IsDigit(rune(key[i+1]))) {
			continue
		}

		// Identify and skip language codes e.g. en-US, zh-cn, en_US, ZH_cn
		if len(key) > i+6 && r == ' ' && internal.IsLetter(rune(key[i+1])) && internal.IsLetter(rune(key[i+2])) && (key[i+3] == '-' || key[i+3] == '_') && internal.IsLetter(rune(key[i+4])) && internal.IsLetter(rune(key[i+5])) && (key[i+6] == ' ' || key[i+6] == ')' || key[i+6] == ';') {
			// Add the number of runes to skip to the skip count.
			skipCount += 6
			continue
		}

		switch r {
		case ' ', ';', ')', '(', ',', '_', '-', '/':
			continue
		}

		// If result exists, we can append it to the value.
		for _, result := range node.result {
			matched := ua.addMatch(result)

			// If we matched a browser of the highest precedence, we can mark the
			// next set of runes as the version number we want to store.
			//
			// We also reject any version numbers related to Safari since it has a
			// separate key for its version number.
			if (matched && result.Type == internal.MatchBrowser && result.Match != internal.Safari) || (result.Type == internal.MatchVersion && ua.versionIndex == 0) {
				// Clear version buffer if it has old values.
				if ua.versionIndex > 0 {
					ua.version = [32]rune{}
					ua.versionIndex = 0
				}

				// We want to omit the slash after the browser name.
				skipCount = 1
				isVersion = true
			}

			// If we matched a mobile token, we want to strip everything after it
			// until we reach whitespace to get around random device IDs.
			// For example, "Mobile/14F89" should be "Mobile".
			if matched && result.Match == internal.Mobile {
				skipUntilWhitespace = true
			}

			// If we matched an Android token, we want to strip everything after it until
			// we reach a closing parenthesis to get around random device IDs.
			if matched && result.Match == internal.Android {
				skipUntilClosingParenthesis = true
			}
		}

		// Set the next node to the child of the current node.
		var next *RuneTrie
		if len(node.childrenArr) != 0 {
			for _, child := range node.childrenArr {
				if child.r == r {
					next = child.node
					break
				}
			}
		} else {
			next = node.childrenMap[r]
		}

		if next == nil {
			continue // No match found, but we can try to match the next rune.
		}
		node = next
	}

	return ua
}

// Put inserts the value into the trie at the given key, replacing any
// existing items. At the end of key tokens, a result is stored marking
// a potential match for a browser, device, or OS using the indexes provided
// by MatchTokenIndexes.
func (trie *RuneTrie) Put(key string) {
	node := trie
	matchResults := internal.MatchTokenIndexes(key)
	for keyIndex, r := range key {
		// Initialise a new result slice for each new rune.
		if node.result == nil {
			node.result = []resultItem{}
		}

		// If we encounter a match, we can store it in the trie.
		for _, result := range matchResults {
			if keyIndex == result.EndIndex-1 {
				newResult := resultItem{Match: result.Match, Type: result.MatchType, Precedence: result.Precedence}
				if !slices.Contains(node.result, newResult) {
					node.result = append(node.result, newResult)
				}
			}
		}

		var child *RuneTrie
		// If the number of children is less than the array size, we can store
		// the children in an array.
		if node.childrenMap == nil && len(node.childrenArr) < maxChildArraySize {
			// Search for the child in the array
			for _, c := range node.childrenArr {
				// If the child is found, set the child to the node.
				if c.r == r {
					child = c.node
					break
				}
			}

			if child == nil {
				// No child found, create a new one
				child = new(RuneTrie)
				node.childrenArr = append(node.childrenArr, childNode{r: r, node: child})
			}
		} else {
			// If the number of children is greater than the array size, we can store
			// the children in a map. We also empty the array.
			if node.childrenMap == nil {
				node.childrenMap = make(map[rune]*RuneTrie)

				// Transfer children from array to map
				for _, c := range node.childrenArr {
					node.childrenMap[c.r] = c.node
				}
				node.childrenArr = nil // Clear the array as it's no longer needed
			}

			// Use the map to store the child
			child = node.childrenMap[r]
			if child == nil {
				child = new(RuneTrie)
				node.childrenMap[r] = child
			}
		}

		node = child
	}
}

// This adds a matching constant to a user agent struct.
func (ua *UserAgent) addMatch(result resultItem) bool {
	ua.Device = "Unknown"
	// Browsers
	if result.Type == internal.MatchBrowser && result.Precedence > ua.browserPrecedence {
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
			ua.device = internal.DeviceMobile
			ua.Device = "Mobile"
		case internal.Safari:
			ua.browser = internal.Safari
		case internal.Vivaldi:
			ua.browser = internal.Vivaldi
		case internal.Samsung:
			ua.browser = internal.Samsung
		case internal.Falkon:
			ua.browser = internal.Falkon
		case internal.Nintendo:
			ua.browser = internal.Nintendo
		case internal.YandexBrowser:
			ua.browser = internal.YandexBrowser
		}

		ua.browserPrecedence = result.Precedence
		return true
	}

	// Operating Systems
	if result.Type == internal.MatchOS && result.Precedence > ua.osPrecedence {
		switch result.Match {
		case internal.Android:
			ua.os = internal.Android
			ua.device = internal.DeviceMobile
			ua.Device = "Mobile"
			// An older generic white-labeled variant of Chrome/Chromium on Android.
			if ua.browser == "" {
				ua.browser = internal.AndroidBrowser
				// Special case we set this as the precedence with this is zero
				// and can be overwritten by Safari.
				ua.browserPrecedence = internal.MatchPrecedenceMap[internal.Mobile]
			}
		case internal.ChromeOS:
			ua.os = internal.ChromeOS
			ua.device = internal.DeviceDesktop
			ua.Device = "Desktop"

		case internal.IOS:
			ua.os = internal.IOS
			if ua.device != internal.DeviceTablet {
				ua.device = internal.DeviceMobile
				ua.Device = "Mobile"
			}
		case internal.Linux:
			ua.os = internal.Linux
			if ua.device != internal.DeviceTablet && ua.device != internal.DeviceTV {
				ua.device = internal.DeviceDesktop
				ua.Device = "Desktop"
			}
		case internal.OpenBSD:
			ua.os = internal.OpenBSD
			ua.device = internal.DeviceDesktop
			ua.Device = "Desktop"
		case internal.MacOS:
			ua.os = internal.MacOS
			ua.device = internal.DeviceDesktop
			ua.Device = "Desktop"
		case internal.Windows:
			ua.os = internal.Windows
			ua.device = internal.DeviceDesktop
			ua.Device = "Desktop"
		}

		ua.osPrecedence = result.Precedence
		return true
	}

	// Types
	if result.Type == internal.MatchType && result.Precedence > ua.typePrecedence {
		switch result.Match {
		case internal.Desktop:
			ua.device = internal.DeviceDesktop
			ua.Device = "Desktop"
		case internal.Tablet:
			ua.device = internal.DeviceTablet
			ua.Device = "Tablet"
		case internal.Mobile, internal.MobileDevice:
			if ua.device != internal.DeviceTablet {
				ua.device = internal.DeviceMobile
				ua.Device = "Mobile"
			}
		case internal.TV:
			ua.device = internal.DeviceTV
			ua.Device = "TV"
		case internal.Bot:
			ua.device = internal.DeviceBot
			ua.Device = "Bot"
		}

		ua.typePrecedence = result.Precedence
		return true
	}

	return false
}
