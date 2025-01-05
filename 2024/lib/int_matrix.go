package lib

import (
	"strings"
)

type IntMatrix struct {
	width  int
	height int
	data   [][]int
}

func NewIntMatrix(width, height int) *IntMatrix {
	data := make([][]int, height)
	for y := 0; y < height; y++ {
		data[y] = make([]int, width)
	}
	return &IntMatrix{
		width:  width,
		height: height,
		data:   data,
	}
}

func NewIntMatrixFromString(input string) *IntMatrix {
	lines := strings.Split(input, lf)
	r := NewIntMatrix(len(lines[0]), len(lines))
	for y, line := range lines {
		for x, char := range ArrStrToInt(strings.Split(line, "")) {
			r.Set(x, y, int(char))
		}
	}
	return r
}

func (m *IntMatrix) Exists(x, y int) bool {
	return x > -1 && x < m.width && y > -1 && y < m.height
}
func (m *IntMatrix) Set(x, y, value int) {
	m.data[y][x] = value
}
func (m *IntMatrix) Get(x, y int) int {
	if !m.Exists(x, y) {
		return -1
	}
	return m.data[y][x]
}

func (r *IntMatrix) String() string {
	lines := make([]string, r.height)
	for y := 0; y < r.height; y++ {
		lines[y] = ArrIntToStr(r.data[y])
	}
	return strings.Join(lines, lf)
}
