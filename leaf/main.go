package main

import (
	"fmt"
	"os"
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
	case "run":
		mainRun(args)
	case "build":
		mainBuild(args)
	case "help":
		mainHelp(args)
	default:
		fmt.Fprintf(os.Stderr, "leaf: unknown subcommand %q.\n", cmd)
		fmt.Fprintf(os.Stderr, "Run 'leaf help' for usage.\n")
	}
}

func mainRun(args []string) {
	panic("todo")
}

func mainHelp(args []string) {
	panic("todo")
}

func intro() {
	panic("todo")
}
