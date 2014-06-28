package main

import (
	"fmt"
	"os"
)

func printErrors(errs []error) {
	if len(errs) > 0 {
		for _, e := range errs {
			printError(e)
		}
	}
}

func printError(e error) {
	fmt.Fprintln(os.Stderr, e)
}
