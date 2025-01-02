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

func getNextDirection(direction string) string {
	switch direction {
	case "up":
		return "right"
	case "right":
		return "down"
	case "down":
		return "left"
	case "left":
		return "up"
	}
	panic("unknown direction")
}

func processInputPartOne(matrix *lib.StringMatrix) {
	fmt.Println(matrix.String())

	x, y := matrix.Find(GUARD_UP)
	if x == -1 || y == -1 {
		panic("could not find start position")
	}

	fmt.Printf("Start: %d:%d\n", x, y)

	direction := "up"
	i := 0
	for matrix.Exists(x, y) && i < 4096 {
		x, y = matrix.WalkInDirection(x, y, DIRECTIONS[direction], WALL)
		direction = getNextDirection(direction)
		if x == -1 && y == -1 {
			fmt.Printf("Exit found at %d:%d\n", x, y)
		} else {
			fmt.Printf("Wall at: %d:%d\n", x, y)
		}
		i++
	}
}
