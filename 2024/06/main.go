package main

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/Ephigenia/advent-of-code/2024/lib"
)

const WALL = "#"
const EMPTY = "."
const GUARD_UP = "^"

var DIRECTIONS = map[string][]int{
	"up":    {0, -1}, // up
	"down":  {0, 1},  // down
	"right": {1, 0},  // right
	"left":  {-1, 0}, // left
}

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

	parsedInput := parseInput(rawInput)

	processInputPartOne(parsedInput)
}

func parseInput(str string) *lib.StringMatrix {
	matrix := lib.NewStringMatrixFromString(str)
	return matrix
}

func processInputPartOne(matrix *lib.StringMatrix) {
	fmt.Println(matrix.String())

	x, y := matrix.Find(GUARD_UP)
	if x == -1 || y == -1 {
		panic("could not find start position")
	}

	fmt.Printf("Start: %d:%d", x, y)

	direction := DIRECTIONS["up"]
	nextHindernis := matrix.FindInDirection(x, y, direction, WALL)

}
