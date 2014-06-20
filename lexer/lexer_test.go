package lexer_test

import (
	"strings"
	"testing"

	"e8vm.net/leaf/lexer"
	. "e8vm.net/leaf/lexer/token"
)

type r struct {
	t   Token
	lit string
}

func TestLexer(t *testing.T) {
	_o := func(s string, exp ...*r) *lexer.Lexer {
		lex := lexer.New(strings.NewReader(s), "")
		i := 0
		for lex.Scan() {
			tok := lex.Token()
			if tok.Token == EOF {
				if i != len(exp) {
					t.Errorf("lex %q: unexpect EOF %d", i)
				}
				continue
			}

			if i >= len(exp) ||
				exp[i].t != tok.Token ||
				exp[i].lit != tok.Lit {

				t.Errorf("lex %q: #%d: %q(%s)",
					s, i, tok.Lit, tok.Token,
				)
			}
			i++
		}

		if i != len(exp) {
			t.Errorf("lex %q: ntoken exp %d, got %d", s, len(exp), i)
		}

		if lex.ScanErr() != nil {
			t.Errorf("lex %q: got scan error %s", s, lex.Err())
		}

		return lex
	}
	o := func(s string, exp ...*r) {
		lex := _o(s, exp...)
		e := lex.Err()
		if e != nil {
			t.Errorf("lex %q: got lex error %s", s, e)
		}
	}
	oe := func(s string, exp ...*r) {
		lex := _o(s, exp...)
		if lex.Err() == nil {
			t.Errorf("lex %q: should be illegal", s)
		}
	}

	m := func(t Token, lit string) *r { return &r{t, lit} }
	n := func(t Token) *r { return &r{t, t.String()} }
	id := func(s string) *r { return &r{Ident, s} }
	sc := n(Semi)
	// eof := n(EOF)

	for _, t := range Keywords() {
		o(t.String()+";", n(t), sc)
	}

	for _, t := range Operators() {
		o(t.String()+";", n(t), sc)
	}

	o("")
	o("   ")
	o("  =", n(Assign))
	o("  =    \t", n(Assign))
	o("=", n(Assign))
	o("==", n(Eq))
	o("3.24 = 3", m(Float, "3.24"), n(Assign), m(Int, "3"), sc)
	o("fun", id("fun"), sc)
	o("0X3 >= 0334", m(Int, "0X3"), n(Geq), m(Int, "0334"), sc)

	// numbers
	o("0", m(Int, "0"), sc)
	o("0.", m(Float, "0."), sc)
	o("9.", m(Float, "9."), sc)
	o("0.7", m(Float, "0.7"), sc)
	o("3.e5", m(Float, "3.e5"), sc)
	o("3e5", m(Float, "3e5"), sc)
	o("3E5", m(Float, "3E5"), sc)
	o("3D5", m(Int, "3"), m(Ident, "D5"), sc)
	oe("3e", m(Int, "3e"), sc)
	oe("3.e", m(Int, "3.e"), sc)
	o(".3", m(Float, ".3"), sc)
	oe(".3e", m(Int, ".3e"), sc)
	oe(".3ef", m(Int, ".3e"), id("f"), sc)
	oe("f.3e", id("f"), m(Int, ".3e"), sc)
	o(".357e-32", m(Float, ".357e-32"), sc)
	o("0.357e-32", m(Float, "0.357e-32"), sc)
	o("03.357e-32", m(Int, "03"), m(Float, ".357e-32"), sc)
	o("3.357e+32", m(Float, "3.357e+32"), sc)
	o(".7e5", m(Float, ".7e5"), sc)
	o("``", m(String, "``"), sc)
	o(`""`, m(String, `""`), sc)
	o(`"string"`, m(String, `"string"`), sc)
	oe(`"str`, m(String, `"str`), sc)
	oe("\"string\n\"", m(String, `"string`), sc, m(String, `"`), sc)

	o("a3", id("a3"), sc)
	o("_A3.come()", id("_A3"), n(Dot),
		id("come"), n(Lparen), n(Rparen),
		sc,
	)
	o("A3(); B3(); C3();",
		id("A3"), n(Lparen), n(Rparen), sc,
		id("B3"), n(Lparen), n(Rparen), sc,
		id("C3"), n(Lparen), n(Rparen), sc,
	)
	oe("$", m(Illegal, "$"))

	// comments
	o("//", m(Comment, "//"))
	o("// something", m(Comment, "// something"))
	o("   /* some */ ", m(Comment, "/* some */"))
	o("   /* some ***/", m(Comment, "/* some ***/"))
	oe("   /* some ***", m(Comment, "/* some ***"))
	o("a3/* some */func", id("a3"), m(Comment, "/* some */"), n(Func))

	// char literals
	o(`' '`, m(Char, `' '`), sc)
	o(`'\''`, m(Char, `'\''`), sc)
	oe(`  ' \''`, m(Char, `' \''`), sc)
	o(`'\n'`, m(Char, `'\n'`), sc)
	o(`'\032'`, m(Char, `'\032'`), sc)
	o(`'\327'`, m(Char, `'\327'`), sc)
	oe(`'\328'`, m(Char, `'\328'`), sc)
	o(`'\x3a'`, m(Char, `'\x3a'`), sc)
	o(`'\xa3'`, m(Char, `'\xa3'`), sc)
	oe(`'\xa'`, m(Char, `'\xa'`), sc)
	oe(`'\xaaa'`, m(Char, `'\xaaa'`), sc)
	o(`'永'`, m(Char, `'永'`), sc)
	oe(`'\ax3'`, m(Char, `'\ax3'`), sc)
	oe(`'\32a'`, m(Char, `'\32a'`), sc)
	oe(`'''`, m(Char, `''`), m(Char, `'`), sc)
}
