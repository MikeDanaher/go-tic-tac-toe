package tic_tac_toe

import (
	"bytes"
)

type MockWriter struct{}

func NewMockWriter() *bytes.Buffer {
	return bytes.NewBuffer(make([]byte, 10))
}
