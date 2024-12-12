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

var directions = [][]int{
	{0, -1},  // up
	{1, -1},  // up right
	{1, 0},   // right
	{1, 1},   // bottom right
	{0, 1},   // bottom
	{-1, 1},  // bottom left
	{-1, 0},  // left
	{-1, -1}, // top left
}

func processInputPartOne(input string) {
	matrix := lib.NewStringMatrixFromString(input)
	// visited := lib.NewRuneMapFromString(input)

	fmt.Print(matrix.String())

	for y, row := range matrix.GetData() {
		for x := range row {
			cur := matrix.Get(x, y)
			if cur == "X" {
				found := matrix.GetInDirection(x, y, directions[2])
				fmt.Printf("found X at %d/%d (%s), %s\n", x, y, cur, cur+found)
			}
		}
	}

	// go through each letter and find XMAS
}
