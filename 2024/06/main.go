package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path"

	"github.com/Ephigenia/advent-of-code/2024/lib"
)

const WALL = "#"
const EMPTY = "."
const GUARD_UP = "^"
const VISITED = "X"

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
	slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.Debug(matrix.String())

	visited := matrix.Clone()
	slog.Debug(visited.String())

	// find the start position
	x, y := matrix.Find(GUARD_UP)
	if x == -1 || y == -1 {
		panic("could not find start position")
	}
	slog.Info(fmt.Sprintf("Start: %d:%d\n", x, y))

	// walk into directions until a wall is hit and then
	// change the direction until the matrix is left
	direction := "up"
	i := 0
	const MAX_ITERATIONS = 4096
	for matrix.Exists(x, y) && i < MAX_ITERATIONS {
		i++
		oldX, oldY := x, y
		// continue walking until next wall is found
		x, y = matrix.WalkInDirection(x, y, DIRECTIONS[direction], WALL)
		// mark the visited positions
		visited.Fill(x, y, oldX, oldY, VISITED)
		// change direction
		direction = getNextDirection(direction)
		slog.Debug(fmt.Sprintf("Walked to: %d:%d %d:%d\n", oldX, oldY, x, y))
		if x != -1 && y != -1 {
			slog.Debug(fmt.Sprintf("Wall at: %d:%d\n", x, y))
		}
	}

	slog.Info(visited.String())
	slog.Info(fmt.Sprintf("Exit found at %d:%d\n", x, y))

}
