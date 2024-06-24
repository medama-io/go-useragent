package useragent

import (
	"fmt"
	"slices"
	"strings"
)

type Result struct {
	Match string
	// 0: Unknown, 1: Browser, 2: OS, 3: Type
	Type uint8
	// Precedence value for each result type to determine which result
	// should be overwritten.
	Precedence uint8
}

// RuneTrie is a trie of runes with string keys and interface{} values.
type RuneTrie struct {
	children map[rune]*RuneTrie
	result   []Result
}

// NewRuneTrie allocates and returns a new *RuneTrie.
func NewRuneTrie() *RuneTrie {
	return new(RuneTrie)
}

// Get returns the value stored at the given key. Returns nil for internal
// nodes or for nodes with a value of nil.
func (trie *RuneTrie) Get(key string) UserAgent {
	node := trie
	ua := UserAgent{
		precedence: Precedence{},
	}

	// Flag to indicate if we are currently iterating over a version number.
	var isVersion bool
	// A buffer to store the version number.
	var versionBuffer strings.Builder
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
			if !IsDigit(r) && r != '.' {
				isVersion = false
			} else {
				// Add to rune buffer
				versionBuffer.WriteRune(r)
				continue
			}
		}

		// We want to strip any other version numbers from other products to get more hits
		// to the trie.
		//
		// We also do not use a switch here as Go does not generate a jump table for
		// switch statements with no integral constants. Benchmarking shows that ops
		// go down if we try to migrate statements like this to a switch.
		if IsDigit(r) || (r == '.' && len(key) > i+1 && IsDigit(rune(key[i+1]))) {
			continue
		}

		// Identify and skip language codes e.g. en-US, zh-cn, en_US, ZH_cn
		if len(key) > i+6 && r == ' ' && IsLetter(rune(key[i+1])) && IsLetter(rune(key[i+2])) && (key[i+3] == '-' || key[i+3] == '_') && IsLetter(rune(key[i+4])) && IsLetter(rune(key[i+5])) && (key[i+6] == ' ' || key[i+6] == ')' || key[i+6] == ';') {
			// Add the number of runes to skip to the skip count.
			skipCount += 6
			continue
		}

		switch r {
		case ' ', ';', ')', '(', ',', '_', '-', '/':
			continue
		}

		// If result exists, we can append it to the value.
		// fmt.Printf(string(r))
		// fmt.Printf("%+v\n", node.result)
		for _, result := range node.result {
			matched := ua.addMatch(result)

			// If we matched a browser of the highest precedence, we can mark the
			// next set of runes as the version number we want to store.
			//
			// We also reject any version numbers related to Safari since it has a
			// separate key for its version number.
			if (matched && result.Type == BrowserMatch && result.Match != Safari) || (result.Type == VersionMatch && ua.Version == "") {
				// Clear version buffer if it has old values.
				versionBuffer.Reset()
				skipCount = 1 // We want to omit the slash after the browser name.
				isVersion = true
			}

			// If we matched a mobile token, we want to strip everything after it
			// until we reach whitespace to get around random device IDs.
			// For example, "Mobile/14F89" should be "Mobile".
			if matched && result.Match == Mobile {
				// We need to clear the result so we can match the next token.
				// node.result = nil
				skipUntilWhitespace = true
			}

			// If we matched an Android token, we want to strip everything after it until
			// we reach a closing parenthesis to get around random device IDs.
			if matched && result.Match == Android {
				// node.result = nil
				skipUntilClosingParenthesis = true
			}

			// fmt.Printf("- %t - ", matched)
			// fmt.Printf("%+v\n", result)
		}

		// We need to catch the flag change after the loop since it isn't possible
		// for a continue to affect an outer loop.
		/* if skipUntilWhitespace {
			continue
		}*/

		// Set the next node to the child of the current node.

		next := node.children[r]
		if next == nil {
			continue // No match found, but we can try to match the next rune.
		}
		node = next
	}
	fmt.Println()

	// Store version buffer into the user agent struct.
	ua.Version = versionBuffer.String()

	return ua
}

// Put inserts the value into the trie at the given key, replacing any
// existing items. At the end of key tokens, a result is stored marking
// a potential match for a browser, device, or OS using the indexes provided
// by MatchTokenIndexes.
func (trie *RuneTrie) Put(key string) {
	node := trie
	matchResults := MatchTokenIndexes(key)
	for keyIndex, r := range key {
		// Initialise a new result slice for each new rune.
		if node.result == nil {
			node.result = []Result{}
		}

		// If we encounter a match, we can store it in the trie.
		for _, result := range matchResults {
			if keyIndex == result.EndIndex-1 {
				newResult := Result{Match: result.Match, Type: result.MatchType, Precedence: result.Precedence}
				if !slices.Contains(node.result, newResult) {
					node.result = append(node.result, newResult)
				}
			}
		}

		// fmt.Printf(string(r))
		// fmt.Printf("%+v\n", node.result)

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
