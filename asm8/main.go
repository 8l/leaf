package main

import (
	"flag"
	"fmt"
	"os"

	"e8vm.net/leaf/asm8/lexer"
	"e8vm.net/leaf/asm8/parser"
)

func main() {
	args := os.Args

	if len(args) <= 1 {
		intro()
		return
	}

	cmd := args[1]
	args = args[2:]

	switch cmd {
	case "lex":
		mainLex(args)
	case "parse":
		mainParse(args)
	default:
		panic("todo")
	}
}

func intro() {
	panic("todo")
}

func mainLex(args []string) {
	fset := flag.NewFlagSet("asm8-lex", flag.ExitOnError)
	fset.Parse(args)

	files := fset.Args()

	if len(files) != 1 {
		fmt.Fprintln(os.Stderr, "expects one single input file.")
		os.Exit(1)
	}

	onError := func(e error) { fmt.Fprintln(os.Stderr, e) }

	f := files[0]
	fin, e := os.Open(f)
	if e != nil {
		onError(e)
		os.Exit(1)
	}

	lex := lexer.New(fin, f)
	lex.OnError(onError)

	for lex.Scan() {
		tok := lex.Token()
		fmt.Println(tok.String())
	}

	e = fin.Close()
	if e != nil {
		onError(e)
	}
}

func mainParse(args []string) {
	fset := flag.NewFlagSet("asm8-parse", flag.ExitOnError)
	fset.Parse(args)

	files := fset.Args()

	if len(files) != 1 {
		fmt.Fprintln(os.Stderr, "expects one single input file.")
		os.Exit(1)
	}

	onError := func(e error) { fmt.Fprintln(os.Stderr, e) }

	f := files[0]
	fin, e := os.Open(f)
	if e != nil {
		onError(e)
		os.Exit(1)
	}

	parser := parser.New(fin, f)
	_, errs := parser.Parse()

	parser.PrintTree(os.Stdout)
	for _, e := range errs {
		onError(e)
	}

	if len(errs) > 0 {
		os.Exit(1)
	}
}
