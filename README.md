# go-useragent

`go-useragent` is a high-performance Go library designed to parse browser name and version, operating system, and device type information from user-agent strings with _sub-microsecond_ parsing times.

It achieves this efficiency by using a modified hybrid [trie](https://en.wikipedia.org/wiki/Trie) data structure to store and rapidly look up user-agent tokens. Utilizing heuristic rules, the library tokenizes a list of user-agent strings into a trie during startup. Subsequently, during runtime, the parsing process involves a straightforward lookup operation.

## Installation

```bash
go get -u github.com/medama-io/go-useragent
```

## Usage

This type of parser is typically initialized once at application startup and reused throughout the application's lifecycle. While it doesn't offer the exhaustive coverage of regex-based parsers, it can be paired with a regex-based parser to handle unknown edge cases, where the trie-based parser acts as a fast path for the majority of user-agents.

## Example

```go
package main

import (
	"fmt"
	"github.com/medama-io/go-useragent"
)

func main() {
	// Create a new parser. Initialize only once during application startup.
	ua := useragent.NewParser()

	// Example user-agent string.
	str := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36"

	// Parse the user-agent string.
	agent := ua.Parse(str)

	// Access parsed information using agent fields.
	fmt.Println(agent.Browser)  // Chrome
	fmt.Println(agent.OS)       // Windows
	fmt.Println(agent.Version)  // 118.0.0.0
	fmt.Println(agent.Desktop)  // true
	fmt.Println(agent.Mobile)   // false
	fmt.Println(agent.Tablet)   // false
	fmt.Println(agent.TV)       // false
	fmt.Println(agent.Bot)      // false

	// Helper functions.
	fmt.Println(agent.GetMajorVersion())  // 118
}
```
Refer to the [pkg.go.dev](https://pkg.go.dev/github.com/medama-io/go-useragent) documentation for more details on available fields and their meanings.

## Acknowledgements

- The library draws inspiration from the techniques outlined in this [Raygun blog post](https://raygun.com/blog/possibility-tree-fast-string-parsing/).