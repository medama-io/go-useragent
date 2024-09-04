package useragent

import (
	"strings"
)

type Parser struct {
	Trie *RuneTrie
}

type UserAgent struct {
	// Precedence is the order in which the user agent matched the
	// browser, device, and OS. The lower the number, the higher the
	// precedence.
	BrowserPrecedence uint8
	OSPrecedence      uint8
	TypePrecedence    uint8

	Browser string
	OS      string
	Version string

	Desktop bool
	Mobile  bool
	Tablet  bool
	TV      bool
	Bot     bool
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
