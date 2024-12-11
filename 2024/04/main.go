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

	processInputPartOne(rawInput)
}

func processInputPartOne(input string) {
	matrix := lib.NewStringMatrixFromString(input)
	// visited := lib.NewRuneMapFromString(input)

	fmt.Print(matrix.String())

	fmt.Printf("%v", matrix.Get(1, 1))

	// go through each letter and find XMAS
}
