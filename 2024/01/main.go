package main

import (
	"errors"
	"fmt"
	"math"
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

	processInputPartOne(rawInput)
}

func processInputPartOne(input string) {
	lines := strings.Split(input, "\n")
	right := make([]int, len(lines))
	left := make([]int, len(lines))

	for i, line := range lines {
		values := lib.ArrStrToInt(strings.Split(line, "   "))
		left[i] = values[0]
		right[i] = values[1]
	}

	// find min values of remaining array and calculate distance
	sum := 0
	for i := range left {
		minLeft, minLeftIndex := lib.ArrMin(left)
		minRight, minRightIndex := lib.ArrMin(right)
		left = lib.ArrIntPopIndex(left, minLeftIndex)
		right = lib.ArrIntPopIndex(right, minRightIndex)

		dist := math.Abs(float64(minLeft - minRight))
		fmt.Printf("#%d minLeft: %d, minRight: %d, dist: %f\n", i, minLeft, minRight, dist)
		sum += int(dist)
	}
	spew.Dump(sum)
}
