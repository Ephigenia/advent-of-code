package main

import (
	"errors"
	"fmt"
	"math"
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
	// processInputPartTwo(rawInput)
}

func RedNoseFactorIsSafe(factor int) bool {
	return factor <= 3
}

func isSafeLevel(level int, nextLevel int) bool {
	delta := int(math.Abs(float64(nextLevel - level)))
	if !RedNoseFactorIsSafe(delta) {
		return false
	}
	return true
}

func processLevels(levels []int, allowedInvalids int) bool {
	deltas := lib.ArrIntDeltas(levels)
	firstSign := 1
	if deltas[0] < 0 {
		firstSign = -1
	}

	for i, delta := range deltas {
		if delta == 0 { // no delta
			fmt.Printf("   #%d invalid because delta 0\n", i)
			return false
		}

		deltaAbs := int(math.Abs(float64(delta)))
		if delta > 0 && firstSign != 1 ||
			delta < 0 && firstSign != -1 { // change of positive to negative
			fmt.Printf("   #%d sign change\n", i)
			return false
		}

		if deltaAbs > 3 {
			fmt.Printf("   #%d invalid delta %d\n", i, deltaAbs)
			return false
		}
		fmt.Printf("   #%d valid delta %d\n", i, deltaAbs)
	}

	return true
}

func processInputPartOne(input string) {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		levels := lib.ArrStrToInt(strings.Split(line, " "))
		result := processLevels(levels, 0)
		fmt.Printf("levels: %v, result: %t\n", levels, result)
		if result {
			sum++
		}
	}
	fmt.Printf("Number of safe reports: %d\n", sum)
}

// func processInputPartTwo(input string) {
// 	lines := strings.Split(input, "\n")

// 	numOfSafeReports := 0
// 	for i, line := range lines {
// 		levels := lib.ArrStrToInt(strings.Split(line, " "))
// 	}
// }
