// Package runes provides some handy functions for rune processing.
package runes

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

func IsDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

func IsOctDigit(r rune) bool {
	return '0' <= r && r <= '7'
}

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

func IsWhite(r rune) bool {
	return r == ' ' || r == '\r' || r == '\t'
}
