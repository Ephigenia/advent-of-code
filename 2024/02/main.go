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
	processInputPartTwo(rawInput)
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

func analyzeDelta(delta int, deltaAbs int, firstSign int) (bool, error) {
	if delta == 0 { // no delta
		return false, errors.New("delta 0")
	} else if delta > 0 && firstSign != 1 ||
		delta < 0 && firstSign != -1 { // change of positive to negative
		return false, errors.New("sign change")
	} else if deltaAbs > 3 {
		return false, errors.New("to large delta")
	}
	return true, nil
}

func processLevels(levels []int, allowedInvalids int) bool {
	errorCount := 0
	deltas := lib.ArrIntDeltas(levels)
	firstSign := 1
	if deltas[0] < 0 {
		firstSign = -1
	}

	for i, delta := range deltas {
		deltaAbs := int(math.Abs(float64(delta)))
		_, err := analyzeDelta(delta, deltaAbs, firstSign)
		if err != nil {
			errorCount++
		}
		if errorCount > allowedInvalids {
			return false
		}
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

func processInputPartTwo(input string) {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		levels := lib.ArrStrToInt(strings.Split(line, " "))
		result := processLevels(levels, 2)
		fmt.Printf("levels: %v, result: %t\n", levels, result)
		if result {
			sum++
		}
	}
	fmt.Printf("Number of safe reports: %d\n", sum)
}
