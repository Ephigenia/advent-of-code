package lib

import (
	"strings"
)

var DIRECTIONS = map[string][]int{
	"up":    {0, -1}, // up
	"down":  {0, 1},  // down
	"right": {1, 0},  // right
	"left":  {-1, 0}, // left
}

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
func (m *IntMatrix) GetWidth() int {
	return m.width
}
func (m *IntMatrix) GetHeight() int {
	return m.height
}

func (r *IntMatrix) String() string {
	lines := make([]string, r.height)
	for y := 0; y < r.height; y++ {
		lines[y] = ArrIntToStr(r.data[y])
	}
	return strings.Join(lines, lf)
}

func (r *IntMatrix) FindAll(v int) [][]int {
	found := make([][]int, 0)
	for y := 0; y < r.height; y++ {
		for x := 0; x < r.width; x++ {
			if r.Get(x, y) == v {
				found = append(found, []int{x, y})
			}
		}
	}
	return found
}

func (m IntMatrix) MovePosition(x, y int, direction string) (int, int) {
	dir := DIRECTIONS[direction]
	return x + dir[0], y + dir[1]
}
