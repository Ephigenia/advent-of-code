package main

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/Ephigenia/advent-of-code/2024/lib"
)

func main() {
	pwd, _ := os.Getwd()
	args := os.Args

	if len(args) < 2 {
		lib.ExitWithError(errors.New("please provide a filename"))
	}
	filename := path.Join(pwd, args[1])

	rawInput, err := lib.ReadInputFile(filename)
	if err != nil {
		lib.ExitWithError(err)
	}
	partOne(rawInput)
}

type Point struct {
	x, y int
}
type Area struct {
	points []Point
}

func (a Area) Width() int {
	w := 0
	for _, p := range a.points {
		if p.x > w {
			w = p.x
		}
	}
	return w + 1
}
func (a Area) Height() int {
	h := 0
	for _, p := range a.points {
		if p.y > h {
			h = p.y
		}
	}
	return h + 1
}

func partOne(input string) {
	matrix := lib.NewStringMatrixFromString(input)
	fmt.Printf("%s\n", matrix.String())

	// parse all areas
	area := Area{points: []Point{
		{x: 0, y: 0},
		{x: 0, y: 1},
		{x: 1, y: 2},
		{x: 1, y: 2},
	}}

	fmt.Printf("w/h: %d/%d\n", area.Width(), area.Height())

}
