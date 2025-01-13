package useragent

import (
	"strings"
	"sync"
)

type device uint8

const (
	deviceUnknown device = iota
	deviceDesktop
	deviceMobile
	deviceTablet
	deviceTV
	deviceBot
)

var (
	once   sync.Once
	parser *Parser
)

type Parser struct {
	Trie *RuneTrie
}

type UserAgent struct {
	browser string
	os      string

	version      [32]rune
	versionIndex int

	// Precedence is the order in which the user agent matched the
	// browser, device, and OS. The lower the number, the higher the
	// precedence.
	browserPrecedence uint8
	osPrecedence      uint8
	typePrecedence    uint8

	device device
}

// Create a new Trie and populate it with user agent data.
func NewParser() *Parser {
	once.Do(func() {
		trie := NewRuneTrie()
		parser = &Parser{Trie: trie}

		// For each newline in the file, add the user agent to the trie.
		for _, ua := range strings.Split(userAgentsFile, "\n") {
			parser.Trie.Put(ua)
		}
	})

	return parser
}

// Parse a user agent string and return a UserAgent struct.
func (p *Parser) Parse(ua string) UserAgent {
	return p.Trie.Get(ua)
}
