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
	return x > -1 && x < r.width && y > -1 && y < r.height
}

func (r *StringMatrix) GetInDirection(x, y int, direction []int) string {
	found := []string{}
	for r.Exists(x, y) {
		x += direction[0]
		y += direction[1]
		c := r.Get(x, y)
		found = append(found, c)
	}

	return strings.Join(found, "")
}

func (r *StringMatrix) FindInDirection(x, y int, direction []int, str string) (int, int) {
	for r.Exists(x, y) {
		// fmt.Printf("found: %d:%d %s\n", x, y, r.Get(x, y))
		if r.Get(x, y) == str {
			return x, y
		}
		x += direction[0]
		y += direction[1]
	}
	return -1, -1
}

func (r *StringMatrix) WalkInDirection(x, y int, direction []int, str string) (int, int) {
	fx, fy := r.FindInDirection(x, y, direction, str)
	if fx > -1 && fy > -1 {
		return fx - direction[0], fy - direction[1]
	}
	return fx, fy
}

func (r *StringMatrix) Find(query string) (int, int) {
	for y := 0; y < r.height; y++ {
		for x := 0; x < r.width; x++ {
			if r.Get(x, y) == query {
				return x, y
			}
		}
	}
	return -1, -1
}
