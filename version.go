package useragent

import (
	"os"
	"strings"
	"unicode"
)

// RemoveVersions removes the version numbers from the user agent string.
func RemoveVersions(ua string) string {
	// Flag to indicate if we are currently iterating over a version number.
	isVersion := false
	isMacVersion := false
	skipCount := uint8(0)
	indexesToReplace := []int{}

	for i, r := range ua {
		if skipCount > 0 {
			skipCount--
			continue
		}

		// If we encouter a slash, we can assume the version number is next.
		if r == '/' {
			isVersion = true
		} else if unicode.IsSpace(r) {
			// If we encounter a space, we can assume the version number is over.
			isVersion = false
		}

		// Mac OS X version numbers are separated by "X " followed by a version number
		// with underscores.
		if r == 'X' && len(ua) > i+1 && unicode.IsSpace(rune(ua[i+1])) {
			isMacVersion = true
		} else if r == ')' {
			isMacVersion = false
		}

		// We want to strip any other version numbers from other products to get more hits
		// to the trie.
		if unicode.IsDigit(r) || (r == '.' && len(ua) > i+1 && unicode.IsDigit(rune(ua[i+1]))) {
			indexesToReplace = append(indexesToReplace, i)
			continue
		}

		// If we are currently iterating over a version number, add the index to the
		// list of indexes to replace. We can't remove the rune in this pass as it
		// would change the indexes of the remaining runes.
		if isVersion || isMacVersion {
			indexesToReplace = append(indexesToReplace, i)
			continue
		}

		// Identify and skip language codes e.g. en-US, zh-cn, en_US, ZH_cn
		if len(ua) > i+6 && r == ' ' && unicode.IsLetter(rune(ua[i+1])) && unicode.IsLetter(rune(ua[i+2])) && (ua[i+3] == '-' || ua[i+3] == '_') && unicode.IsLetter(rune(ua[i+4])) && unicode.IsLetter(rune(ua[i+5])) && (ua[i+6] == ' ' || ua[i+6] == ')' || ua[i+6] == ';') {
			// Add the number of runes to skip to the skip count.
			skipCount += 6
			indexesToReplace = append(indexesToReplace, i, i+1, i+2, i+3, i+4, i+5, i+6)
			continue
		}

		// Skip whitespace
		if r == ' ' || r == ';' || r == ')' || r == '(' || r == ',' || r == '_' {
			indexesToReplace = append(indexesToReplace, i)
			continue
		}
	}

	// Remove the version numbers from the user agent string.
	for _, i := range indexesToReplace {
		ua = ua[:i] + ua[i+1:]
		// Update the indexes of the remaining runes.
		for j := range indexesToReplace {
			if indexesToReplace[j] > i {
				indexesToReplace[j]--
			}
		}
	}

	return ua
}

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
		cleanedLine := RemoveVersions(line)

		// For each line, get all token indexes
		// and remove all strings after the largest EndIndex.
		results := MatchTokenIndexes(cleanedLine)

		// If no results, skip the line.
		if len(results) == 0 {
			continue
		}

		// Get the largest EndIndex.
		largestEndIndex := results[0].EndIndex
		// Remove all strings after the largest EndIndex.
		cleanedLine = cleanedLine[:largestEndIndex]

		// Check for duplicates.
		if !seen[cleanedLine] {
			cleanedAgents = append(cleanedAgents, cleanedLine)
			seen[cleanedLine] = true
		}
	}

	return cleanedAgents, nil
}
