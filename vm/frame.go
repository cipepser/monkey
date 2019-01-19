package vm

import (
	"github.com/cipepser/monkey/code"
	"github.com/cipepser/monkey/object"
)

type Frame struct {
	cl         *object.Closure
	ip         int
	basePinter int
}

func NewFrame(cl *object.Closure, basePointer int) *Frame {
	return &Frame{
		cl:         cl,
		ip:         -1,
		basePinter: basePointer,
	}
}

func (f *Frame) Instructions() code.Instructions {
	return f.cl.Fn.Instructions
}
