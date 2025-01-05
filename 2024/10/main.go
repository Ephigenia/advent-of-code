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

func partOne(matrix *lib.IntMatrix) {
	// find starting positions
	startPositions := matrix.FindAll(0)
	spew.Dump(startPositions)
}
