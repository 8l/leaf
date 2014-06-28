.PHONY: check build lex parse

run:
	leaf build *.lf
	e8 out.e8

check:
	leaf lex main.lf | diff - lex.result
	leaf parse main.lf | diff - parse.result

build:
	leaf lex main.lf > lex.result
	leaf parse main.lf > parse.result

lex:
	leaf lex main.lf

parse:
	leaf parse main.lf

ast:
	leaf parse -ast main.lf
