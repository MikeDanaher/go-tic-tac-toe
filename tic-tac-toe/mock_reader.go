package tic_tac_toe

import (
	"bytes"
)

type MockReader struct{}

func NewMockReader(s string) *bytes.Buffer {
	return bytes.NewBufferString(s)
}
