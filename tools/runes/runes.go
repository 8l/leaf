// Package runes provides some handy functions for rune processing.
package runes

// IsLetter tests if r is a letter (a-z and A-Z) or an underscore.
func IsLetter(r rune) bool {
	if 'a' <= r && r <= 'z' {
		return true
	}
	if 'A' <= r && r <= 'Z' {
		return true
	}
	if r == '_' {
		return true
	}
	return false
}

// IsDigit tests if r is a digit (0-9).
func IsDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

// IsOctDigit tests if r is an octal digit (0-7).
func IsOctDigit(r rune) bool {
	return '0' <= r && r <= '7'
}

// IsHexDigit tests if r is a heximal digit (0-9, a-f or a-F).
func IsHexDigit(r rune) bool {
	if IsDigit(r) {
		return true
	}
	if 'A' <= r && r <= 'F' {
		return true
	}
	if 'a' <= r && r <= 'f' {
		return true
	}
	return false
}

// IsWhite tests if r is a space, '\r' or '\t'. It returns false for '\n'.
func IsWhite(r rune) bool {
	return r == ' ' || r == '\r' || r == '\t'
}
