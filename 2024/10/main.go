package main

import (
	"errors"
	"fmt"
	"os"
	"path"

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

func trackback(matrix *lib.IntMatrix, x, y int) {
	for direction := range lib.DIRECTIONS {
		cur := matrix.Get(x, y)
		x, y := matrix.MovePosition(x, y, direction)
		if !matrix.Exists(x, y) {
			continue
		}
		next := matrix.Get(x, y)
		diff := next - cur
		if diff != 1 {
			continue
		}
		fmt.Printf("found %d at %d,%d %s\n", next, x, y, direction)
	}
}

func partOne(matrix *lib.IntMatrix) {
	// find starting positions
	startPositions := matrix.FindAll(0)
	// iterate over all start positions and find trails
	for posI, pos := range startPositions {
		fmt.Printf("#%d start at %v\n", posI, pos)
		trackback(matrix, pos[0], pos[1])
	}
	spew.Dump(startPositions)
}
