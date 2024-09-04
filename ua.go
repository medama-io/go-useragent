package useragent

import (
	"strings"
)

type Parser struct {
	Trie *RuneTrie
}

// Precedence is the order in which the user agent matched the
// browser, device, and OS. The lower the number, the higher the
// precedence.
type Precedence struct {
	Browser uint8
	OS      uint8
	Type    uint8
}

type UserAgent struct {
	Browser string
	OS      string
	Version string

	Desktop bool
	Mobile  bool
	Tablet  bool
	TV      bool
	Bot     bool

	precedence Precedence
}

// Create a new Trie and populate it with user agent data.
func NewParser() *Parser {
	trie := NewRuneTrie()
	parser := &Parser{Trie: trie}

	// For each newline in the file, add the user agent to the trie.
	for _, ua := range strings.Split(userAgentsFile, "\n") {
		parser.Trie.Put(ua)
	}

	return parser
}

// Parse a user agent string and return a UserAgent struct.
func (p *Parser) Parse(ua string) UserAgent {
	return p.Trie.Get(ua)
}
