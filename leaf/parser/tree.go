package parser

import (
	"io"
)

type TokenTree interface {
	PrintTree(io.Writer)
}

func ParseTree(f string) (TokenTree, []error) {
	p, e := Open(f)
	if e != nil {
		return nil, []error{e}
	}

	_, errs := p.Parse()
	return p, errs
}

func (p *Parser) PrintTree(out io.Writer) {
	p.s.PrintTrack(out)
}
