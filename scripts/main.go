package main

import (
	"fmt"
	"os"
	"strings"

	ua "github.com/medama-io/go-useragent"
)

func main() {
	var content []string
	filenames := []string{"agents/1.txt", "agents/2.txt"}

	for _, filename := range filenames {
		// Read agents.txt file.
		agents, err := ua.CleanAgentsFile(filename)
		if err != nil {
			fmt.Printf("Error cleaning agents file: %s\n", err)
			return
		}

		content = append(content, agents...)
	}

	// Check for duplicates.
	seen := make(map[string]bool) // to track duplicates
	var contentNoDuplicates []string
	for _, line := range content {
		if !seen[line] {
			contentNoDuplicates = append(contentNoDuplicates, line)
			seen[line] = true
		}
	}

	// Join the cleaned agents into a single string.
	finalStr := strings.Join(contentNoDuplicates, "\n")

	// Write the cleaned content to agents_cleaned.txt file.
	writePath := "agents/final.txt"
	err := os.WriteFile(writePath, []byte(finalStr), 0o644)
	if err != nil {
		fmt.Printf("Error writing cleaned agents file: %s\n", err)
		return
	}

	fmt.Println("Cleaned agents saved to agents_cleaned.txt")
}
