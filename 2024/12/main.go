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

func partOne(input string) {
	matrix := lib.NewStringMatrixFromString(input)
	fmt.Printf("%s\n", matrix.String())

	// areas := []lib.StringMatrix{}
	matrix.ForEach(func(x, y int, val string) {
		fmt.Printf("x/y: %d/%d\n", x, y)
	})
}
