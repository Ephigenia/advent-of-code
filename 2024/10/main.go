package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/Ephigenia/advent-of-code/2024/lib"
	"github.com/davecgh/go-spew/spew"
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

	matrix := lib.NewIntMatrixFromString(rawInput)
	fmt.Printf("%v\n", matrix)
	partOne(matrix)
}

func trackback(matrix *lib.IntMatrix, x, y int) [][]int {
	found := make([][]int, 0)
	visited := matrix.Clone()

	for direction := range lib.DIRECTIONS {
		cur := matrix.Get(x, y)
		x, y := matrix.MovePosition(x, y, direction)

		if !matrix.Exists(x, y) {
			continue
		}
		if cur == 9 {
			continue
		}

		if visited.Get(x, y) == 1 {
			continue
		}
		visited.Set(x, y, 1)

		if matrix.Get(x, y)-cur != 1 {
			continue
		}

		found = append(found, trackback(matrix, x, y)...)
		found = append(found, []int{x, y})
	}
	return found
}

func formatPath(path [][]int) string {
	parts := make([]string, len(path))
	for i, pos := range path {
		parts[i] = fmt.Sprintf("%d:%d", pos[0], pos[1])
	}
	return strings.Join(parts, " -> ")
}

func partOne(matrix *lib.IntMatrix) {
	// find starting positions
	startPositions := matrix.FindAll(0)
	// iterate over all start positions and find trails
	for posI, pos := range startPositions {
		path := trackback(matrix, pos[0], pos[1])
		fmt.Printf("#%d start at %v: %s\n", posI, pos, formatPath(path))
	}
	spew.Dump(startPositions)
}
