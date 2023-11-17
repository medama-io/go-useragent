package useragent

import (
	"os"
	"strings"
)

// ReplaceIndexes replaces the runes at the given indexes with empty strings.
func ReplaceIndexes(ua string, indexes []int) string {
	// Remove the version numbers from the user agent string.
	for _, i := range indexes {
		ua = ua[:i] + ua[i+1:]
		// Update the indexes of the remaining runes.
		for j := range indexes {
			if indexes[j] > i {
				indexes[j]--
			}
		}
	}

	return ua
}

// RemoveVersions removes the version numbers from the user agent string.
func RemoveVersions(ua string) string {
	// Flag to indicate if we are currently iterating over a version number.
	var isVersion bool
	// Number of runes to skip when iterating over the trie. This is used
	// to skip over version numbers or language codes.
	var skipCount uint8
	// Indexes of the runes to replace with an empty string.
	indexesToReplace := []int{}

	for i, r := range ua {
		if skipCount > 0 {
			skipCount--
			continue
		}

		if isVersion {
			// If we encounter any unknown characters, we can assume the version number is over.]
			if !IsDigit(r) && r != '.' {
				isVersion = false
			} else {
				// If we are currently iterating over a version number, add the index to the
				// list of indexes to replace. We can't remove the rune in this pass as it
				// would change the indexes of the remaining runes.
				indexesToReplace = append(indexesToReplace, i)
				continue
			}
		}

		// We want to strip any other version numbers from other products to get more hits
		// to the trie.
		if IsDigit(r) || (r == '.' && len(ua) > i+1 && IsDigit(rune(ua[i+1]))) {
			indexesToReplace = append(indexesToReplace, i)
			continue
		}

		// Identify and skip language codes e.g. en-US, zh-cn, en_US, ZH_cn
		if len(ua) > i+6 && r == ' ' && IsLetter(rune(ua[i+1])) && IsLetter(rune(ua[i+2])) && (ua[i+3] == '-' || ua[i+3] == '_') && IsLetter(rune(ua[i+4])) && IsLetter(rune(ua[i+5])) && (ua[i+6] == ' ' || ua[i+6] == ')' || ua[i+6] == ';') {
			// Add the number of runes to skip to the skip count.
			skipCount += 6
			indexesToReplace = append(indexesToReplace, i, i+1, i+2, i+3, i+4, i+5, i+6)
			continue
		}

		// Skip whitespace
		switch r {
		case ' ', ';', ')', '(', ',', '_', '-', '/':
			indexesToReplace = append(indexesToReplace, i)
			continue
		}

		// Replace all non-latin characters with a space. The trie function will automatically
		// skip over any characters it can't find, so this is a safe operation.
		if !IsLetter(r) {
			indexesToReplace = append(indexesToReplace, i)
			continue
		}
	}

	ua = ReplaceIndexes(ua, indexesToReplace)
	return ua
}

// RemoveDeviceIdentifiers removes the device identifiers from the user agent string.
// This specifically removes any strings that follow the Mobile tokens.
func RemoveDeviceIdentifiers(ua string, tokens []MatchResults) string {
	// Find mobile token.
	for _, token := range tokens {
		var skipUntilWhitespace bool
		var indexesToReplace []int
		if token.Match == Mobile {
			// Iterate over the user agent string and remove all characters
			// after the mobile token until we encounter whitespace.
			for i, r := range ua {
				if skipUntilWhitespace {
					if r == ' ' {
						skipUntilWhitespace = false
					} else {
						indexesToReplace = append(indexesToReplace, i)
						continue
					}
				}

				if i == token.EndIndex-1 {
					skipUntilWhitespace = true
				}
			}

			ua = ReplaceIndexes(ua, indexesToReplace)
			return ua
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
		line = RemoveDeviceIdentifiers(line, MatchTokenIndexes(line))
		line = RemoveVersions(line)

		// For each line, get all token indexes
		// and remove all strings after the largest EndIndex.
		results := MatchTokenIndexes(line)

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
