package lib

import (
	"fmt"
	"strings"
)

type Grid struct {
	Cells         [][]rune
	Width, Height int
}

func NewGrid(width, height int) *Grid {
	cells := make([][]rune, height)
	for i := range cells {
		cells[i] = make([]rune, width)
	}
	return &Grid{
		Cells:  cells,
		Width:  width,
		Height: height,
	}
}

func NewGridFromString(input string) *Grid {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	height := len(lines)
	width := len(lines[0])
	grid := NewGrid(width, height)
	for y, line := range lines {
		for x, char := range line {
			grid.Set(x, y, rune(char))
		}
	}
	return grid
}

func (g *Grid) Get(x, y int) rune {
	return g.Cells[y][x]
}

func (g *Grid) Exists(x, y int) bool {
	return x >= 0 && x < g.Width && y >= 0 && y < g.Height
}

func (g *Grid) Set(x, y int, value rune) {
	g.Cells[y][x] = value
}

func (g *Grid) Print() {
	for _, row := range g.Cells {
		for _, cell := range row {
			fmt.Printf("%s", string(cell))
		}
		fmt.Println()
	}
}

type Offset struct {
	DX, DY int
}

func (g *Grid) GetOffsets(x, y int, offsets []Offset) []rune {
	ret := make([]rune, len(offsets))
	for i, offset := range offsets {
		nx, ny := x+offset.DX, y+offset.DY
		ret[i] = -1
		if g.Exists(nx, ny) {
			ret[i] = g.Get(nx, ny)
		}
	}
	return ret
}

// returns the runes around the given coordinates
func (g *Grid) GetS(x, y int) []rune {
	offsets := []Offset{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
	return g.GetOffsets(x, y, offsets)
}
