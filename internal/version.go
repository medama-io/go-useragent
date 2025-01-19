package internal

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

// RemoveMobileIdentifiers removes the device identifiers from the user agent string.
// This specifically removes any strings that follow the Mobile tokens.
// For example, "Mobile/14F89" should be "Mobile".
func RemoveMobileIdentifiers(ua string) string {
	tokens := MatchTokenIndexes(ua)

	// Find mobile token.
	for _, token := range tokens {
		var skipUntilWhitespace bool
		var indexesToReplace []int
		if token.Match == DeviceMobile {
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

// RemoveAndroidIdentifiers removes the device identifiers from the user agent string.
// This specifically removes any strings that follow the Android tokens.
func RemoveAndroidIdentifiers(ua string) string {
	tokens := MatchTokenIndexes(ua)

	// Find mobile token.
	for _, token := range tokens {
		var skipUntilClosingParenthesis int
		var indexesToReplace []int

		if token.Match == OSAndroid {
			// Iterate over the user agent string and remove all characters
			// after the Android token until we encounter a closing parenthesis
			// to remove device identifiers.
			for i, r := range ua {
				if skipUntilClosingParenthesis > 0 {
					if r == '(' {
						skipUntilClosingParenthesis++
					}

					if r == ')' {
						skipUntilClosingParenthesis--
					} else {
						indexesToReplace = append(indexesToReplace, i)
						continue
					}
				}

				if i == token.EndIndex-1 {
					skipUntilClosingParenthesis++
				}
			}

			ua = ReplaceIndexes(ua, indexesToReplace)
			return ua
		}
	}

	return ua
}
