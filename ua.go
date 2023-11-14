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

	Desktop bool
	Mobile  bool
	Tablet  bool

	precedence Precedence
}

// populateTrie populates the trie with user agent data.
func (p *Parser) populateTrie() {
	p.trie.Put("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36")
}

// Create a new Trie and populate it with user agent data.
func NewParser() *Parser {
	trie := NewRuneTrie()
	parser := &Parser{trie: trie}
	parser.populateTrie()
	return parser
}

// Parse a user agent string and return a UserAgent struct.
func (p *Parser) Parse(ua string) *UserAgent {
	return p.trie.Get(ua)
}
