package tic_tac_toe

import (
	"io"
)

type Input struct {
	reader io.Reader
}

func NewInput(reader io.Reader) *Input {
	return &Input{reader}
}

func (i *Input) Read() string {
	buf := make([]byte, 10)
	n, _ := i.reader.Read(buf)
	text := string(buf[:n])
	return text
}
