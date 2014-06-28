.PHONY: check build lex parse

run:
	leaf build *.lf
	e8 out.e8

lex:
	leaf lex main.lf

parse:
	leaf parse main.lf

ast:
	leaf parse -ast main.lf
