package tic_tac_toe

import (
	"io"
)

type Output struct {
	writer io.Writer
}

func NewOutput(writer io.Writer) *Output {
	return &Output{writer}
}

func (o *Output) Write(s string) {
	p := []byte(s)
	o.writer.Write(p)
}
