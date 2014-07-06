package main

import (
	"flag"
	"fmt"
	"os"

	"e8vm.net/leaf/ast"
	"e8vm.net/leaf/parser"
	"e8vm.net/util/prt"
)

func mainParse(args []string) {
	fset := flag.NewFlagSet("leaf-parse", flag.ExitOnError)
	astFlag := fset.Bool("ast", false, "print AST instead of token tree")

	fset.Parse(args)

	files := fset.Args()

	if len(files) == 0 {
		fmt.Fprintln(os.Stderr, "no input file.")
		return
	}

	for _, f := range files {
		fmt.Printf("[%s]\n", f)

		if *astFlag {
			res, errs := parser.Parse(f)
			printErrors(errs)
			if res != nil {
				p := prt.New(os.Stdout)
				p.Indent = "    "
				ast.Print(p, res)
			}
		} else {
			tree, errs := parser.ParseTree(f)
			printErrors(errs)
			if tree != nil { // might be nil when the file does not exist
				tree.PrintTree(os.Stdout)
			}
		}
	}
}
