package useragent

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
	result   *Result
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
	var isVersion, isMacVersion bool
	// Number of runes to skip when iterating over the trie. This is used
	// to skip over version numbers or language codes.
	var skipCount uint8

	for i, r := range key {
		if skipCount > 0 {
			skipCount--
			continue
		}

		// If we encounter a potential version, skip the runes until we reach
		// the end of the version number.
		switch r {
		case '/':
			isVersion = true
		case ' ':
			// If we encounter a space, we can assume the version number is over.
			isVersion = false
		}

		// Mac OS X version numbers are separated by "X " followed by a version number
		// with underscores.
		//
		// We also do not use a switch here as Go does not generate a jump table for
		// switch statements with no integral constants. Benchmarking shows that ops
		// go down if we try to migrate statements like this to a switch.
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

		// Identify and skip language codes e.g. en-US, zh-cn, en_US, ZH_cn
		if len(key) > i+6 && r == ' ' && IsLetter(rune(key[i+1])) && IsLetter(rune(key[i+2])) && (key[i+3] == '-' || key[i+3] == '_') && IsLetter(rune(key[i+4])) && IsLetter(rune(key[i+5])) && (key[i+6] == ' ' || key[i+6] == ')' || key[i+6] == ';') {
			// Add the number of runes to skip to the skip count.
			skipCount += 6
			continue
		}

		switch r {
		case ' ', ';', ')', '(', ',', '_', '-':
			continue
		}

		// If result exists, we can append it to the value.
		if node.result != nil {
			ua.addMatch(node.result)
		}

		// Set the next node to the child of the current node.
		next := node.children[r]
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
	matchResults := MatchTokenIndexes(key)
	for i, r := range key {
		// If we've reached the end of a matching key, store the result.
		matchIndex := len(matchResults) - 1
		// The end index is after the last rune in the match, so
		// we need to subtract 1 to get the last rune.
		if i == matchResults[matchIndex].EndIndex-1 {
			node.result = &Result{Match: matchResults[matchIndex].Match, Type: matchResults[matchIndex].MatchType, Precedence: matchResults[matchIndex].Precedence}
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
