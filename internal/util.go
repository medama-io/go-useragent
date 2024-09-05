package internal

// IsDigit reports whether the rune is a decimal digit.
//
// This is an optimised version of unicode.IsDigit without
// the check for r <= unicode.MaxLatin1.
func IsDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

// IsLetter reports whether the rune is a letter in ASCII.
//
// This is an optimised version of unicode.IsLetter without
// the check for r <= unicode.MaxLatin1.
func IsLetter(r rune) bool {
	return 'a' <= r && r <= 'z' || 'A' <= r && r <= 'Z'
}
