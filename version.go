package useragent

// RemoveVersions removes the version numbers from the user agent string.
func RemoveVersions(ua string) string {
	// Flag to indicate if we are currently iterating over a version number.
	isVersion := false
	isMacVersion := false
	indexesToReplace := []int{}

	for i, r := range ua {
		// If we encouter a slash, we can assume the version number is next.
		if r == '/' {
			isVersion = true
		} else if r == ' ' {
			// If we encounter a space, we can assume the version number is over.
			isVersion = false
		}

		// Mac OS X version numbers are separated by "X " followed by a version number
		// with underscores.
		if r == 'X' && ua[i+1] == ' ' {
			isMacVersion = true
		} else if r == ')' {
			isMacVersion = false
		}

		// If we are currently iterating over a version number, add the index to the
		// list of indexes to replace. We can't remove the rune in this pass as it
		// would change the indexes of the remaining runes.
		if isVersion || isMacVersion {
			indexesToReplace = append(indexesToReplace, i)
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
