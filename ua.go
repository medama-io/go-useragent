package useragent

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/medama-io/go-useragent/internal"
)

var (
	once   sync.Once
	parser *Parser
)

type Parser struct {
	Trie *RuneTrie
}

type UserAgent struct {
	version      [32]rune
	versionIndex int

	browser internal.Match
	os      internal.Match
	device  internal.Match

	// Precedence is the order in which the user agent matched the
	// browser, device, and OS. The lower the number, the higher the
	// precedence.
	browserPrecedence uint8
	osPrecedence      uint8
	typePrecedence    uint8
}

// Parse a user agent string and return a UserAgent struct.
func (p *Parser) Parse(ua string) UserAgent {
	return p.Trie.Get(ua)
}

// NewParser creates a new parser and populates it with the default embedded user agent data.
func NewParser() *Parser {
	once.Do(func() {
		var err error

		parser, err = newParserFromReader(strings.NewReader(userAgentsFile))
		if err != nil {
			// Panicking is fine since it would be caught in a test and is a fixed trusted input.
			panic("failed to parse embedded user agent definitions: " + err.Error())
		}

		// For each newline in the file, add the user agent to the trie.
		for _, ua := range strings.Split(userAgentsFile, "\n") {
			parser.Trie.Put(ua)
		}
	})

	return parser
}

// NewParserWithFile creates a new parser with user agent definitions loaded from a file.
//
// The file should contain one user agent definition per line.
func NewParserWithFile(filePath string) (*Parser, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filePath, err)
	}
	defer file.Close()

	return newParserFromReader(file)
}

// NewParserWithURL creates a new parser with user agent definitions loaded from a URL.
//
// The URL should serve content with one user agent definition per line.
func NewParserWithURL(url string) (*Parser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status %d: %s", resp.StatusCode, resp.Status)
	}

	return newParserFromReader(resp.Body)
}

func newParserFromReader(reader io.Reader) (*Parser, error) {
	trie := NewRuneTrie()
	parser := &Parser{Trie: trie}

	scanner := bufio.NewScanner(reader)
	lineCount := 0

	for scanner.Scan() {
		parser.Trie.Put(scanner.Text())
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading user agent definitions: %w", err)
	}

	if lineCount == 0 {
		return nil, errors.New("no user agent definitions found")
	}

	return parser, nil
}
