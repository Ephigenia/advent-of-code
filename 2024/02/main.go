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

	processInputPartOne(rawInput)
}

func RedNoseFactor(one, two int) int {
	// return int(math.Abs(float64(one - two)))
	return one - two
}

func RedNoseFactorIsSafe(factor int) bool {
	return factor <= 3
}

func processInputPartOne(input string) {
	lines := strings.Split(input, "\n")

	for i, line := range lines {
		levels := lib.ArrStrToInt(strings.Split(line, " "))

		deltas := lib.ArrIntMap(levels, func(v int, i int) int {
			if i == len(levels)-1 {
				return 0
			}
			return RedNoseFactor(v, levels[i+1])
		})
		deltasAbs := lib.ArrIntAbs(deltas)

		deltaSigns := lib.ArrIntMap(deltas, func(v int, i int) int {
			if v < 0 {
				return -1
			}
			return 1
		})
		deltaSignsSum := lib.ArrIntSum(deltaSigns)

		max, _ := lib.ArrMax(deltas)
		isSafe := deltaSignsSum == len(deltas)
		if max > 3 {
			isSafe = false
		}

		fmt.Printf("#%d levels: %v, factors %v (%v) isSafe: %t\n", i, levels, deltas, deltasAbs, isSafe)
	}
}
