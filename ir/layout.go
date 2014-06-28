package ir

import (
	"e8vm.net/e8/mem"
)

const (
	segSize = mem.SegSize

	ioStart    = mem.SegIO
	codeStart  = mem.SegCode
	heapStart  = mem.SegHeap
	stackStart = mem.SegStack
)
