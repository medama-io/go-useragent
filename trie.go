package useragent

import "unicode"

type Result struct {
	result string
	// 0: Unknown, 1: Browser, 2: OS, 3: Type
	resultType uint8
	// Precedence value for each result type to determine which result
	// should be overwritten.
	precedence uint8
}

// RuneTrie is a trie of runes with string keys and interface{} values.
type RuneTrie struct {
	children map[rune]*RuneTrie
	result   *Result
	// Number of runes to skip when iterating over the trie. This is used
	// to skip over version numbers or language codes.
	skipIndex int
}

// NewRuneTrie allocates and returns a new *RuneTrie.
func NewRuneTrie() *RuneTrie {
	return new(RuneTrie)
}

// IsDigit reports whether the rune is a decimal digit.
func IsDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

// Get returns the value stored at the given key. Returns nil for internal
// nodes or for nodes with a value of nil.
func (trie *RuneTrie) Get(key string) *UserAgent {
	node := trie
	ua := &UserAgent{
		precedence: Precedence{},
	}

	// Flag to indicate if we are currently iterating over a version number.
	isVersion := false
	isMacVersion := false
	skipCount := uint8(0)

	for i, r := range key {
		if skipCount > 0 {
			skipCount--
			continue
		}

		// If we encounter a potential version, skip the runes until we reach
		// the end of the version number.
		if r == '/' {
			isVersion = true
			continue
		} else if r == ' ' {
			isVersion = false
		}

		// Mac OS X version numbers are separated by "X " followed by a version number
		// with underscores.
		if r == 'X' && len(key) > i+1 && key[i+1] == ' ' {
			isMacVersion = true
		} else if r == ')' {
			isMacVersion = false
		}

		if isVersion || isMacVersion {
			continue
		}

		// We want to strip any other version numbers from other products to get more hits
		// to the trie.
		if IsDigit(r) || (r == '.' && len(key) > i+1 && IsDigit(rune(key[i+1]))) {
			continue
		}

		// Identify and skip language codes e.g. en-US, zh-cn
		if len(key) > i+6 && r == ' ' && unicode.IsLetter(rune(key[i+1])) && unicode.IsLetter(rune(key[i+2])) && (key[i+3] == '-' || key[i+3] == '_') && unicode.IsLetter(rune(key[i+4])) && unicode.IsLetter(rune(key[i+5])) && (key[i+6] == ' ' || key[i+6] == ')' || key[i+6] == ';') {
			// Add the number of runes to skip to the skip count.
			skipCount += 6
			continue
		}

		// If result exists, we can append it to the value.
		if node.result != nil {
			ua.addMatch(node.result, ua.precedence)
		}

		node = node.children[r]
		if node == nil {
			return ua
		}

	}
	return ua
}

// Put inserts the value into the trie at the given key, replacing any
// existing items. At the end of key tokens, a result is stored marking
// a potential match for a browser, device, or OS using the indexes provided
// by MatchTokenIndexes.
func (trie *RuneTrie) Put(key string) {
	node := trie
	key = RemoveVersions(key)
	matchResults := MatchTokenIndexes(key)
	for i, r := range key {
		// If we've reached the end of the key, store the result.
		matchIndex := len(matchResults) - 1
		if matchIndex != -1 && i == matchResults[matchIndex].EndIndex {
			node.result = &Result{result: matchResults[matchIndex].Match, resultType: matchResults[matchIndex].MatchType, precedence: matchResults[matchIndex].Precedence}
			matchResults = matchResults[:matchIndex]
		}

		child := node.children[r]
		if child == nil {
			// If no map is found, create a new one.
			if node.children == nil {
				node.children = map[rune]*RuneTrie{}
			}

			// Store new runes in the trie.
			child = new(RuneTrie)
			node.children[r] = child
		}
		node = child
	}
}
