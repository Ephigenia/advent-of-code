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

	processInputPartTwo(rawInput)
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

func foundAndMas(matrix *lib.StringMatrix, x, y int) bool {
	TL := matrix.Get(x-1, y-1)
	TR := matrix.Get(x+1, y-1)
	BL := matrix.Get(x-1, y+1)
	BR := matrix.Get(x+1, y+1)
	if (TL == "M" && TR == "M" && BL == "S" && BR == "S") ||
		(TL == "S" && TR == "M" && BL == "S" && BR == "M") ||
		(TL == "M" && TR == "S" && BL == "M" && BR == "S") ||
		(TL == "S" && TR == "S" && BL == "M" && BR == "M") {
		return true
	}
	return false
}

func processInputPartTwo(input string) {
	matrix := lib.NewStringMatrixFromString(input)

	fmt.Print(matrix.String() + "\n")

	occurrences := 0

	for y, row := range matrix.GetData() {
		for x := range row {
			cur := matrix.Get(x, y)
			if cur == "A" {
				fmt.Printf("found \"A\" at %d:%d\n", x, y)
				found := foundAndMas(matrix, x, y)
				if found {
					fmt.Printf("found \"A\" at %d:%d <---\n", x, y)
					occurrences++
				}
			}
		}
	}
	fmt.Printf("Found %d occurrences of MAS\n", occurrences)
}

func processInputPartOne(input string) {
	matrix := lib.NewStringMatrixFromString(input)
	// visited := lib.NewRuneMapFromString(input)

	fmt.Print(matrix.String() + "\n")

	occurrences := 0

	for y, row := range matrix.GetData() {
		for x := range row {
			cur := matrix.Get(x, y)
			for _, direction := range directions {
				if cur == "X" {
					isFoundX, isFoundY := matrix.FindInDirection(x, y, direction, "XMAS")
					if isFoundX > -1 && isFoundY > -1 {
						occurrences++
						found := matrix.GetInDirection(x, y, direction)
						fmt.Printf("found \"X\" at %d:%d (%s), %s\n", x, y, cur, cur+found)
					}
				}
			}
		}
	}

	fmt.Printf("Found %d occurrences of XMAS\n", occurrences)
}
