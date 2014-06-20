package runes

import (
	"testing"
)

var symbols = "~!@#$%^&*()-=+[]{}\\|;:'\"<>,.?/`"

func TestRunes(t *testing.T) {
	as := func(c bool) {
		if !c {
			t.Fail()
		}
	}

	asRange := func(from, to rune, f func(rune) bool) {
		for c := from; c <= to; c++ {
			as(f(c))
		}
	}

	asNotRange := func(from, to rune, f func(rune) bool) {
		for c := from; c <= to; c++ {
			as(!f(c))
		}
	}

	asNotIn := func(s string, f func(rune) bool) {
		for _, c := range s {
			as(!f(c))
		}
	}

	asRange('a', 'z', IsLetter)
	asRange('A', 'Z', IsLetter)
	as(IsLetter('_'))
	asNotRange('0', '9', IsLetter)
	asNotIn(symbols, IsLetter)

	asRange('0', '9', IsDigit)
	asRange('0', '9', IsHexDigit)
	asRange('0', '7', IsOctDigit)
	asNotRange('8', '9', IsOctDigit)
	asNotRange('a', 'z', IsOctDigit)
	asNotRange('A', 'Z', IsOctDigit)
	asNotRange('a', 'z', IsDigit)
	asNotRange('A', 'Z', IsDigit)
	asRange('a', 'f', IsHexDigit)
	asRange('A', 'F', IsHexDigit)
	asNotRange('g', 'z', IsHexDigit)
	asNotRange('G', 'Z', IsHexDigit)
	asNotIn(symbols, IsDigit)
	asNotIn(symbols, IsOctDigit)
	asNotIn(symbols, IsHexDigit)

	as(IsWhite('\t'))
	as(IsWhite(' '))
	as(IsWhite('\r'))
	as(!IsWhite('\n'))
}
