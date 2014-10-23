package tic_tac_toe

import (
	"bytes"
	"strings"
)

type Input struct {
	reader *bytes.Buffer
}

func NewInput(reader *bytes.Buffer) *Input {
	return &Input{reader}
}

func (i *Input) Read() string {
	line, _ := i.reader.ReadString('\n')
	return strings.Trim(line, "\n")
}
