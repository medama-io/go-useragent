package useragent

// IsDigit reports whether the rune is a decimal digit.
func IsDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

// IsLetter reports whether the rune is a letter in ASCII.
func IsLetter(r rune) bool {
	return 'a' <= r && r <= 'z' || 'A' <= r && r <= 'Z'
}
