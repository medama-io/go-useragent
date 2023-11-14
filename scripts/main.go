package main

import (
	"fmt"

	ua "github.com/medama-io/go-useragent"
)

func main() {
	err := ua.CleanAgentsFile()
	if err != nil {
		fmt.Printf("Error cleaning agents file: %s\n", err)
	}
}
