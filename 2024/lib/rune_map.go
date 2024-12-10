package lib

import (
	"fmt"
	"strings"
)

const lf = "\n"

type RuneMap struct {
	data   [][]rune
	width  int
	height int
}

func NewRuneMap(width, height int) *RuneMap {
	return &RuneMap{
		data:   make([][]rune, height, width),
		width:  width,
		height: height,
	}
}

func NewRuneMapFromString(input string) *RuneMap {

	lines := strings.Split(input, lf)
	height := len(lines)
	width := len(lines[0])
	r := NewRuneMap(width, height)
	for y, line := range lines {
		for x, char := range line {
			fmt.Printf("x: %d, y: %d, char: %c\n", x, y, char)
			r.Set(x, y, char)
		}
	}
	return r
}

func (r *RuneMap) GetData() [][]rune {
	return r.data
}

func (r *RuneMap) String() string {
	lines := make([]string, r.height)
	for y := 0; y < r.height; y++ {
		lines[y] = string(r.data[y]) + lf
	}
	return strings.Join(lines, lf)
}

func (r *RuneMap) Set(x, y int, value rune) {
	r.data[y][x] = value
}

func (r *RuneMap) Get(x, y int) rune {
	return r.data[y][x]
}
