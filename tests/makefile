.PHONY: check build lex parse

run:
	leaf run *.l

lex:
	leaf lex main.l

parse:
	leaf parse main.l

ast:
	leaf parse -ast main.l
