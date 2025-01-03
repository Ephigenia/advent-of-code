package main

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/Ephigenia/advent-of-code/2024/lib"
	"github.com/davecgh/go-spew/spew"
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
	// find the start position
	x, y := matrix.Find(GUARD_UP)
	if x == -1 || y == -1 {
		panic("could not find start position")
	}
	fmt.Println(fmt.Sprintf("Start: %d:%d\n", x, y))
	matrix.Fill(x, y, x, y, EMPTY) // remove starting margin

	matrix.Fill(7, 7, 7, 10, VISITED)

	// walk into directions until a wall is hit and then
	// change the direction until the matrix is left
	visited := matrix.Clone()
	direction := "left"
	found := true
	i := 0
	const MAX_ITERATIONS = 4096
	for found && i < MAX_ITERATIONS {
		i++
		oldX, oldY := x, y

		// continue walking until next wall is found
		x, y, found = matrix.WalkInDirection(x, y, DIRECTIONS[direction], WALL)

		if found {
			// change direction
			direction = getNextDirection(direction)
			fmt.Println(fmt.Sprintf("Wall at: %d:%d, going %s", x, y, direction))
		}
		// mark the visited positions
		visited.Fill(x, y, oldX, oldY, VISITED)
	}

	fmt.Println(matrix.String() + "\n")
	fmt.Println(visited.String())
	spew.Dump(visited.FindAll(VISITED))
	fmt.Printf("visited nodes: %d\n", len(visited.FindAll(VISITED)))
	fmt.Println(fmt.Sprintf("Exit found at %d:%d", x, y))
}
