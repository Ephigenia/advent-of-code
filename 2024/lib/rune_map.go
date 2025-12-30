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

func (r *StringMatrix) GetWidth() int {
	return r.width
}

func (r *StringMatrix) GetHeight() int {
	return r.height
}

func (r *StringMatrix) Fill(x1, y1, x2, y2 int, str string) {
	if x1 > x2 {
		tmp := x1
		x1 = x2
		x2 = tmp
	}
	if y1 > y2 {
		tmp := y1
		y1 = y2
		y2 = tmp
	}
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			r.Set(x, y, str)
		}
	}
}

func (r *StringMatrix) String() string {
	lines := make([]string, r.height)
	for y := 0; y < r.height; y++ {
		lines[y] = strings.Join(r.data[y], "")
	}
	return strings.Join(lines, lf)
}

func (r *StringMatrix) Set(x, y int, value string) {
	if x == -1 {
		return
	}
	if y == -1 {
		return
	}
	if x >= r.width-1 {
		x = r.width - 1
	}
	if y >= r.width-1 {
		y = r.width - 1
	}
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

func movePosition(x, y int, direction []int) (int, int) {
	return x + direction[0], y + direction[1]
}

func (r *StringMatrix) GetInDirection(x, y int, direction []int) string {
	found := []string{}
	for r.Exists(x, y) {
		x, y = movePosition(x, y, direction)
		c := r.Get(x, y)
		found = append(found, c)
	}

	return strings.Join(found, "")
}

func (r *StringMatrix) FindInDirection(x, y int, direction []int, str string) (int, int, bool) {
	for r.Exists(x, y) {
		if r.Get(x, y) == str {
			return x, y, true
		}
		x, y = movePosition(x, y, direction)
	}
	return x, y, false
}

func (r *StringMatrix) WalkInDirection(x, y int, direction []int, str string) (int, int, bool) {
	fx, fy, found := r.FindInDirection(x, y, direction, str)
	if found {
		return fx - direction[0], fy - direction[1], found
	}
	return fx, fy, found
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

func (r *StringMatrix) Clone() *StringMatrix {
	clone := NewStringMatrix(r.width, r.height)
	for y := 0; y < r.height; y++ {
		for x := 0; x < r.width; x++ {
			clone.Set(x, y, r.Get(x, y))
		}
	}
	return clone
}

func (r *StringMatrix) FindAll(query string) [][]int {
	found := make([][]int, 0)
	for y := 0; y < r.height; y++ {
		for x := 0; x < r.width; x++ {
			if r.Get(x, y) == query {
				found = append(found, []int{x, y})
			}
		}
	}
	return found
}

func (r *StringMatrix) ForEach(fn func(x, y int, value string)) {
	for y := 0; y < r.height; y++ {
		for x := 0; x < r.width; x++ {
			fn(x, y, r.Get(x, y))
		}
	}
}
