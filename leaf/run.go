package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"e8vm.net/leaf/leaf/script"
)

func mainRun(args []string) {

	if len(args) != 1 {
		printError(fmt.Errorf("need one argument to run"))
		os.Exit(1)
	}

	src, e := ioutil.ReadFile(args[0])
	if e != nil {
		printError(e)
		os.Exit(1)
	}

	run := new(script.Run)
	run.Filename = args[0]
	run.Source = src
	run.Stdout = os.Stdout
	run.Run()

	if len(run.Errors) > 0 {
		printErrors(run.Errors)
		os.Exit(1)
	}

	if !run.RIP() {
		fmt.Printf("(ret=%d)\n", run.HaltValue)
		if run.AddrError {
			printError(fmt.Errorf("vm halted on address error"))
		}
	}

	fmt.Printf("(%d cycles)\n", run.UsedCycle)
}
