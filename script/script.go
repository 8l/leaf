// Package script provides APIs to run a single leaf file in an E8 simulator.
package script

import (
	"bytes"
	"io"
	"math"

	"e8vm.net/e8/img"
	"e8vm.net/e8/mem"
	"e8vm.net/leaf/codegen"
	"e8vm.net/leaf/ir"
	"e8vm.net/leaf/parser"
)

type Run struct {
	// input
	Filename     string // file name for the input source
	Source       []byte // the input source file
	Stdout       io.Writer
	TimeoutCycle int

	// output
	Image     []byte  // the compiled E8 image file
	Errors    []error // errors encountered
	Output    []byte  // output of execution
	UsedCycle int
	HaltValue uint8
	AddrError bool
}

func (self *Run) RIP() bool {
	return self.HaltValue == 0 && !self.AddrError
}

func (self *Run) err(es ...error) bool {
	for _, e := range es {
		if e == nil {
			continue
		}
		self.Errors = append(self.Errors, e)
	}
	return len(self.Errors) > 0
}

func (self *Run) Run() {
	if self.Filename == "" {
		self.Filename = "main.l"
	}

	build := ir.NewBuild()
	pname := "main"
	gen := codegen.NewGen(pname, build)
	astree, errs := parser.ParseBytes(self.Filename, self.Source)
	if self.err(errs...) {
		return
	}
	gen.AddFile(astree)

	if self.err(gen.Gen()...) {
		return
	}

	buf := new(bytes.Buffer)
	errs = build.Build(pname, buf) // TODO: stderr?
	if self.err(errs...) {
		return
	}

	self.Image = buf.Bytes()

	vm, e := img.Make(bytes.NewBuffer(self.Image))
	if self.err(e) {
		return
	}

	stdoutBuf := new(bytes.Buffer)
	var stdout io.Writer = stdoutBuf
	if self.Stdout != nil {
		stdout = io.MultiWriter(stdout, self.Stdout)
	}

	vm.SetPC(mem.SegCode)
	vm.Stdout = stdout

	if self.TimeoutCycle == 0 {
		self.TimeoutCycle = math.MaxInt64
	}

	ncycle := vm.Run(self.TimeoutCycle)

	self.UsedCycle = ncycle
	self.HaltValue = vm.HaltValue()
	self.AddrError = vm.AddrError()
	self.Output = stdoutBuf.Bytes()
}
