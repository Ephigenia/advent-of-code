package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

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

func checkStep(cur, prev int) bool {
	if cur == -1 {
		return false
	}
	if cur-prev != 1 {
		return false
	}

	return true
}

func walk(matrix, visited *lib.IntMatrix, prev, x, y int) [][]int {
	path := [][]int{}
	cur := matrix.Get(x, y)

	if !checkStep(cur, prev) {
		return path
	}

	if cur == 9 && prev == 8 {
		if visited.Get(x, y) == 10 {
			return path
		}
		visited.Set(x, y, 10)
		path = append(path, []int{x, y})
		return path
	}

	for direction := range lib.DIRECTIONS {
		x, y := matrix.MovePosition(x, y, direction)
		path = append(path, walk(matrix, visited, cur, x, y)...)
	}
	return path
}

func formatPath(path [][]int) string {
	parts := make([]string, len(path))
	for i, pos := range path {
		parts[i] = fmt.Sprintf("%d:%d ", pos[0], pos[1])
	}
	return strings.Join(parts, ",")
}

func partOne(matrix *lib.IntMatrix) {
	// find starting positions
	startPositions := matrix.FindAll(0)
	// iterate over all start positions and find trails
	sum := 0

	uniquePaths := []string{}

	for posI, pos := range startPositions {
		visited := matrix.Clone()
		path := walk(matrix, visited, -1, pos[0], pos[1])
		fmt.Printf("#%d trail head %d:%d: %d\n", posI, pos[0], pos[1], len(path))

		pathStr := formatPath(path)
		if lib.ArrStrIndexOf(uniquePaths, pathStr) == -1 {
			uniquePaths = append(uniquePaths, pathStr)
		}

		sum += len(path)
	}
	fmt.Printf("\nSum %d\n", sum)
	fmt.Printf("\nUnique Paths: %d\n", len(uniquePaths))
}
