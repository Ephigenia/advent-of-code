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

	matrix := lib.NewIntMatrixFromString(rawInput)
	fmt.Printf("%v\n", matrix)
	partOne(matrix)
}

func walk(matrix, visited *lib.IntMatrix, prev, x, y int) int {
	cur := matrix.Get(x, y)

	if cur == -1 {
		return 0
	}
	if cur-prev != 1 {
		return 0
	}

	if cur == 9 && prev == 8 {
		if visited.Get(x, y) == 10 {
			return 0
		}
		visited.Set(x, y, 10)
		return 1
	}

	ans := 0
	for direction := range lib.DIRECTIONS {
		x, y := matrix.MovePosition(x, y, direction)
		ans += walk(matrix, visited, cur, x, y)
	}
	return ans
}

func partOne(matrix *lib.IntMatrix) {
	// find starting positions
	startPositions := matrix.FindAll(0)
	// iterate over all start positions and find trails
	sum := 0
	for posI, pos := range startPositions {
		visited := matrix.Clone()
		count := walk(matrix, visited, -1, pos[0], pos[1])
		fmt.Printf("#%d trail head %d:%d: %d\n", posI, pos[0], pos[1], count)
		sum += count
	}
	fmt.Printf("Sum %d\n", sum)
}
