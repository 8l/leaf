package main

import (
	"flag"
	"fmt"
	"os"

	"e8vm.net/leaf/asm8/ast"
	"e8vm.net/leaf/asm8/build"
	"e8vm.net/leaf/asm8/codegen"
	"e8vm.net/leaf/asm8/lexer"
	"e8vm.net/leaf/asm8/parser"
	"e8vm.net/leaf/tools/prt"
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
	case "build":
		mainBuild(args)
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
	astFlag := fset.Bool("ast", false, "print AST instead of token tree")
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
		return
	}

	parser := parser.New(fin, f)
	res, errs := parser.Parse()

	if *astFlag {
		if res != nil {
			p := prt.Stdout()
			p.Indent = "    "
			ast.Print(p, res)
		}
	} else {
		parser.PrintTree(os.Stdout)
	}

	for _, e := range errs {
		onError(e)
	}

	if len(errs) > 0 {
		os.Exit(1)
		return
	}
}

func mainBuild(args []string) {
	fset := flag.NewFlagSet("asm8-build", flag.ExitOnError)
	outFlag := fset.String("o", "out.e8", "output e8 image")
	fset.Parse(args)

	files := fset.Args()

	if len(files) != 1 {
		fmt.Fprintln(os.Stderr, "expect one single input file.")
		os.Exit(1)
	}

	onError := func(e error) {
		if e == nil {
			return
		}
		if e != nil {
			fmt.Fprintln(os.Stderr, e)
			os.Exit(1)
			panic("not reachable")
		}
	}

	onErrors := func(es []error) {
		if len(es) == 0 {
			return
		}
		for _, e := range es {
			fmt.Fprintln(os.Stderr, e)
		}
		os.Exit(1)
		panic("not reachable")
	}

	f := files[0]
	fin, e := os.Open(f)
	onError(e)

	parser := parser.New(fin, f)
	astree, errs := parser.Parse()
	onErrors(errs)

	onError(fin.Close())

	b := build.NewBuild()
	gen := codegen.NewGen(b, astree)
	errs = gen.Gen()
	onErrors(errs)

	fout, e := os.Create(*outFlag)
	onError(e)

	e = b.WriteImage(fout)
	onError(e)

	onError(fout.Close())
}
