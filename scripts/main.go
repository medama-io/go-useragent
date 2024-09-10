package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/medama-io/go-useragent/internal"
)

// This reads the agents.txt file and returns a new agents_cleaned.txt file
// with the version numbers removed.
func CleanAgentsFile(filePath string) ([]string, error) {
	// Read agents.txt file.
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Split the content into lines.
	lines := strings.Split(string(content), "\n")

	// Clean each line and store the cleaned agents.
	var cleanedAgents []string
	seen := make(map[string]bool) // to track duplicates
	for _, line := range lines {
		// Check for any invalid lines.
		if len(line) == 0 || len(line) > 400 {
			continue
		}

		if strings.Contains(line, "javascript") || strings.Contains(line, "function") || strings.Contains(line, "quot") || strings.Contains(line, "parent") {
			continue
		}

		line = internal.RemoveMobileIdentifiers(line)
		line = internal.RemoveAndroidIdentifiers(line)
		line = internal.RemoveVersions(line)

		// For each line, get all token indexes
		// and remove all strings after the largest EndIndex.
		results := internal.MatchTokenIndexes(line)

		// If no results, skip the line.
		if len(results) == 0 {
			continue
		}

		// Get the largest EndIndex.
		largestEndIndex := results[0].EndIndex
		// Remove all strings after the largest EndIndex.
		line = line[:largestEndIndex]

		// Check for duplicates.
		if !seen[line] {
			cleanedAgents = append(cleanedAgents, line)
			seen[line] = true
		}
	}

	return cleanedAgents, nil
}

func main() {
	var content []string
	filenames := []string{"agents/1.txt", "agents/2.txt", "agents/3.txt"}

	for _, filename := range filenames {
		// Read agents.txt file.
		agents, err := CleanAgentsFile(filename)
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
