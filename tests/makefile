.PHONY: check build lex parse

run:
	leaf build *.l
	e8 out.e8

lex:
	leaf lex main.l

parse:
	leaf parse main.l

ast:
	leaf parse -ast main.l
