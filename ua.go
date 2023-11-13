package useragent

type Parser struct {
	trie *RuneTrie
}

// Precedence is the order in which the user agent matched the
// browser, device, and OS. The lower the number, the higher the
// precedence.
type Precedence struct {
	Browser int
	OS      int
	Device  int
}

type UserAgent struct {
	Browser string
	OS      string
	Device  string

	precedence *Precedence
}

// Create a new Trie and populate it with user agent data.
func NewParser() *Parser {
	trie := NewRuneTrie()

	return &Parser{trie: trie}
}

// Parse a user agent string and return a UserAgent struct.
func (p *Parser) Parse(ua string) *UserAgent {
	return p.trie.Get(ua)
}
