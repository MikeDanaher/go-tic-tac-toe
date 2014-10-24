package tic_tac_toe

type Reader interface {
	Read(p []byte) (int, error)
}

type Writer interface {
	WriteString(s string) (int, error)
}
