package useragent

type Result struct {
	result     string
	precedence int
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
func (trie *RuneTrie) Get(key string) *UserAgent {
	node := trie
	ua := &UserAgent{
		precedence: &Precedence{},
	}

	for _, r := range key {
		// If result exists, we can append it to the value.
		if node.result != nil {
			ua.addMatch(node.result, ua.precedence)
		}

		node = node.children[r]
		if node == nil {
			return nil
		}

	}
	return ua
}

// Put inserts the value into the trie at the given key, replacing any
// existing items. At the end of key tokens, a result is stored marking
// a potential match for a browser, device, or OS using the indexes provided
// by MatchTokenIndexes.
func (trie *RuneTrie) Put(key string, matchResults []MatchResults) {
	node := trie
	for i, r := range key {
		// If we've reached the end of the key, store the result.
		matchIndex := len(matchResults) - 1
		if matchIndex != -1 && i == matchResults[matchIndex].EndIndex {
			node.result = &Result{result: matchResults[matchIndex].Match, precedence: matchResults[matchIndex].Precedence}
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
