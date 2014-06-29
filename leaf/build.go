package main

import (
	"os"

	"e8vm.net/leaf/codegen"
	"e8vm.net/leaf/ir"
	"e8vm.net/leaf/parser"
)

func mainBuild(args []string) {
	build := ir.NewBuild()

	pname := "test"
	gen := codegen.NewGen(pname, build)
	files := args

	for _, f := range files {
		astree, errs := parser.Parse(f)
		if len(errs) > 0 {
			printErrors(errs)
			os.Exit(-1)
		}

		gen.AddFile(astree)
	}

	errs := gen.Gen()
	if len(errs) > 0 {
		printErrors(errs)
		os.Exit(1)
	}

	// build.Print()

	fout, e := os.Create("out.e8")
	if e != nil {
		printError(e)
		os.Exit(1)
	}

	errs = build.Build(pname, fout)
	if len(errs) > 0 {
		printErrors(errs)
		os.Exit(1)
	}

	e = fout.Close()
	if e != nil {
		printError(e)
		os.Exit(1)
	}
}
