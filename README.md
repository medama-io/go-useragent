# go-useragent

[![Go Reference](https://pkg.go.dev/badge/github.com/medama-io/go-useragent.svg)](https://pkg.go.dev/github.com/medama-io/go-useragent)

`go-useragent` is a high-performance zero-allocation Go library designed to parse browser name and version, operating system, and device type information from user-agent strings with _sub-microsecond_ parsing times.

It achieves this efficiency by using a modified hybrid [trie](https://en.wikipedia.org/wiki/Trie) data structure to store and rapidly look up user-agent tokens. It utilizes heuristic rules, tokenizing a list of user-agent strings into a trie during startup. During runtime, the parsing process involves a straightforward lookup operation.

This project is actively maintained and used by the lightweight website analytics project [Medama](https://github.com/medama-io/medama).

## Installation

```bash
go get -u github.com/medama-io/go-useragent
```

## Usage

This type of parser is typically initialized once at application startup and reused throughout the application's lifecycle. While it doesn't offer the exhaustive coverage of traditional regex-based parsers, it can be paired with one to handle unknown edge cases, where the trie-based parser acts as a fast path for the majority of user-agents.

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
	fmt.Println(agent.Browser())        // agents.BrowserChrome
	fmt.Println(agent.BrowserVersion()) // 118.0.0.0
	fmt.Println(agent.OS())             // agents.OSWindows
	fmt.Println(agent.Device())         // agents.DeviceDesktop

	// Boolean helper functions.
	fmt.Println(agent.IsChrome())  // true
	fmt.Println(agent.IsFirefox()) // false

	fmt.Println(agent.IsWindows()) // true
	fmt.Println(agent.IsLinux())   // false

	fmt.Println(agent.IsDesktop()) // true
	fmt.Println(agent.IsTV())      // false
	fmt.Println(agent.IsBot())     // false
	// and many more...

	// Version helper functions.
	fmt.Println(agent.BrowserVersionMajor()) // 118
	fmt.Println(agent.BrowserVersionMinor()) // 0
	fmt.Println(agent.BrowserVersionPatch()) // 0.0
}
```

Refer to the [pkg.go.dev](https://pkg.go.dev/github.com/medama-io/go-useragent) documentation for more details on available fields and their meanings.

## Benchmarks

Benchmarks were performed against [`ua-parser/uap-go`](https://github.com/ua-parser/uap-go) and [`mileusena/useragent`](https://github.com/mileusna/useragent) on an Apple M3 Pro Processor.

```bash
cd ./benchmarks
go test -bench=. -benchmem ./...

MedamaParserGetSingle-12        3871813             308.3 ns/op               0 B/op          0 allocs/op
MileusnaParserGetSingle-12      1322602             917.3 ns/op             600 B/op         16 allocs/op
UAPParserGetSingle-12            986428              1159 ns/op             233 B/op          8 allocs/op

MedamaParserGetAll-12             71804             15544 ns/op               0 B/op          0 allocs/op
MileusnaParserGetAll-12           28375             42301 ns/op           28031 B/op        716 allocs/op
UAPParserGetAll-12                18645             56951 ns/op           10179 B/op        344 allocs/op
```

## Acknowledgements

- The library draws inspiration from the techniques outlined in this [Raygun blog post](https://raygun.com/blog/possibility-tree-fast-string-parsing/).
