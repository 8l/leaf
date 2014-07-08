package scanner

// NumType indicates what kind of number a token is.
type NumType int

// Enumerates all number types: Illegal, Int, or Float
const (
	NumIllegal NumType = iota
	NumInt
	NumFloat
)

// ScanNumber scans (and accepts) a number using the scanner.
// It returns the literal of the number, and returns
// if the number is a integer, a float or an illegal one.
func ScanNumber(s *Scanner, dotLed bool) (lit string, t NumType) {
	if !dotLed {
		if s.Scan('0') {
			if s.Scan('x') || s.Scan('X') {
				if s.ScanHexDigits() == 0 {
					return s.Accept(), NumIllegal
				}
			} else if s.ScanOctDigit() {
				s.ScanOctDigits()
				return s.Accept(), NumInt
			}

			if s.Peek() != '.' {
				return s.Accept(), NumInt
			}
		}

		s.ScanDigits()

		if s.ScanAny("eE") {
			s.ScanAny("-+")
			if s.ScanDigits() == 0 {
				return s.Accept(), NumIllegal
			}
			return s.Accept(), NumFloat
		}

		if !s.Scan('.') {
			return s.Accept(), NumInt
		}

		s.ScanDigits()
	} else {
		if s.ScanDigits() == 0 {
			return s.Accept(), NumIllegal
		}
	}

	if s.ScanAny("eE") {
		s.ScanAny("-+")
		if s.ScanDigits() == 0 {
			return s.Accept(), NumIllegal
		}
	}

	return s.Accept(), NumFloat
}

// scanEscape scans for an escape rune. and returns any error.
// q is the leading quote rune.
func scanEscape(s *Scanner, q rune) error {
	if s.ScanAny("abfnrtv\\") {
		return nil
	}
	if s.Scan(q) {
		return nil
	}

	if s.Scan('x') {
		if !(s.ScanHexDigit() && s.ScanHexDigit()) {
			return s.Errorf("invalid hex escape")
		}
		return nil
	}

	if s.ScanOctDigit() {
		if !(s.ScanOctDigit() && s.ScanOctDigit()) {
			return s.Errorf("invalid octal escape")
		}
		return nil
	}

	c := s.Peek()
	s.Next()
	return s.Errorf("unknown escape char %q", c)
}

// ScanChar scans for a char, where the leading '\'' is already scanned.
func ScanChar(s *Scanner) (string, error) {
	var e error
	n := 0
	for !s.Scan('\'') {
		if s.Peek() == '\n' || s.Closed() {
			e = s.Errorf("char not terminated")
			break
		}

		if s.Scan('\\') {
			e = scanEscape(s, '\'')
		} else {
			s.Next()
		}
		n++
	}

	if n != 1 && e == nil {
		e = s.Errorf("illegal char")
	}

	return s.Accept(), e
}

// ScanString scans for a string leading by a double-quote '\"',
// where the leading '\"' is already scanned.
func ScanString(s *Scanner) (string, error) {
	var e error
	for !s.Scan('"') {
		if s.Peek() == '\n' || s.Closed() {
			e = s.Errorf("string not terminated")
			break
		}

		if s.Scan('\\') {
			e = scanEscape(s, '"')
		} else {
			s.Next()
		}
	}

	return s.Accept(), e
}

// ScanRawString scans for a string leading by a back-quote.
// A raw string can span multiple lines, and only another back-quote
// can terminate the string. The leading back-quote is already scanned.
func ScanRawString(s *Scanner) (string, error) {
	var e error
	for !s.Scan('`') {
		if s.Closed() {
			e = s.Errorf("raw string not terminated")
			break
		}
		s.Next()
	}
	return s.Accept(), e
}

// ScanComment scans for a comment, where the leading '/' is already scanned.
// The next rune on the scanner must be either '*' or '/'; it will panic
// otherwise.
func ScanComment(s *Scanner) (string, error) {
	var e error

	if s.Scan('*') {
		for {
			if s.Scan('*') {
				if s.Scan('/') {
					return s.Accept(), nil
				}
				continue
			}

			if s.Closed() {
				e = s.Errorf("incomplete block comment")
				return s.Accept(), e
			}
			s.Next()
		}
	}

	if s.Scan('/') {
		for {
			if s.Peek() == '\n' || s.Closed() {
				return s.Accept(), nil
			}
			s.Next()
		}
	}

	panic("unreachable")
}
