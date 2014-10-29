package tic_tac_toe

type Config3x3 struct {
	in  Reader
	out Writer
}

func NewConfig3x3(in Reader, out Writer) Config {
	return &Config3x3{in, out}
}

func (c *Config3x3) BoardSize() int {
	return 3
}

func (c *Config3x3) WinningLines() [][]int {
	return [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
		[]int{0, 3, 6},
		[]int{1, 4, 7},
		[]int{2, 5, 8},
		[]int{0, 4, 8},
		[]int{2, 4, 6},
	}
}

func (c *Config3x3) Input() Reader {
	return c.in
}

func (c *Config3x3) Output() Writer {
	return c.out
}
