# go-useragent

`go-useragent` is a high-performance Go library designed to parse browser, operating system, and device type information from user-agent strings with _sub-microsecond_ parsing times.

It achieves this efficiency by using a [trie](https://en.wikipedia.org/wiki/Trie) data structure to store and rapidly look up user-agent tokens. Utilizing heuristic rules, the library tokenizes a list of user-agent strings into a trie during startup. Subsequently, during runtime, the parsing process involves a straightforward lookup operation with zero allocations.

The project's scope is intentionally limited to the most common user-agents, making it a specialized solution rather than a comprehensive user-agent parser. However, developers can still leverage this library as an initial parser to optimize best-case performance and fallback to a more complete parser to handle the remaining user-agents.

## Installation

```bash
go get -u github.com/medama-io/go-useragent
```

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
	fmt.Println(agent.Desktop)  // true
	fmt.Println(agent.Mobile)   // false
	fmt.Println(agent.Tablet)   // false
	fmt.Println(agent.TV)       // false
	fmt.Println(agent.Bot)      // false
}
```
Refer to the Go.pkg documentation for more details on available fields and their meanings.

## Acknowledgements

- The library draws inspiration from the techniques outlined in this [Raygun blog post](https://raygun.com/blog/possibility-tree-fast-string-parsing/).