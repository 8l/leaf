.PHONY: check build lex parse

run:
	leaf build *.l
	e8 out.e8

check:
	leaf lex main.l | diff - lex.result
	leaf parse main.l | diff - parse.result

build:
	leaf lex main.l > lex.result
	leaf parse main.l > parse.result

lex:
	leaf lex main.l

parse:
	leaf parse main.l

ast:
	leaf parse -ast main.l
