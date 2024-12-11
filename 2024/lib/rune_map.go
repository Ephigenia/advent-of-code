package lib

import (
	"strings"
)

const lf = "\n"

type StringMatrix struct {
	data   [][]string
	width  int
	height int
}

func NewStringMatrix(width, height int) *StringMatrix {
	// make 2d array
	data := make([][]string, height)
	for y := 0; y < height; y++ {
		data[y] = make([]string, width)
	}
	return &StringMatrix{
		data:   data,
		width:  width,
		height: height,
	}
}

func NewStringMatrixFromString(input string) *StringMatrix {
	lines := strings.Split(input, lf)
	height := len(lines)
	width := len(lines[0])
	r := NewStringMatrix(width, height)
	for y, line := range lines {
		for x, char := range line {
			r.Set(x, y, string(char))
		}
	}
	return r
}

func (r *StringMatrix) GetData() [][]string {
	return r.data
}

func (r *StringMatrix) String() string {
	lines := make([]string, r.height)
	for y := 0; y < r.height; y++ {
		lines[y] = strings.Join(r.data[y], "")
	}
	return strings.Join(lines, lf)
}

func (r *StringMatrix) Set(x, y int, value string) {
	r.data[y][x] = value
}

func (r *StringMatrix) Get(x, y int) string {
	if !r.Exists(x, y) {
		return ""
	}
	return r.data[y][x]
}

func (r *StringMatrix) Exists(x, y int) bool {
	return x > 0 && x < r.width && y > 0 && y < r.height
}
