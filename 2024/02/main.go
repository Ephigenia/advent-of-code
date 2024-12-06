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

func RedNoseFactorIsSafe(factor int) bool {
	return factor <= 3
}

func processInputPartOne(input string) {
	lines := strings.Split(input, "\n")

	numOfSafeReports := 0
	for i, line := range lines {
		levels := lib.ArrStrToInt(strings.Split(line, " "))

		// calculate deltas between the level values
		deltas := lib.ArrIntDeltas(levels)
		// find maximum deviation in the deltas
		deltasAbs := lib.ArrIntAbs(deltas)
		maxDelta, _ := lib.ArrMax(deltasAbs)

		// find out if all values are increasing or decreasing
		deltaSigns := lib.ArrIntMap(deltas, func(v int, i int) int {
			if v < 0 {
				return -1
			}
			return 1
		})
		deltaSignsSum := lib.ArrIntSum(deltaSigns)
		isSameDeltaSign := deltaSignsSum == len(deltas) || deltaSignsSum == -len(deltas)

		// increments of 0 are invalid
		containsZeroDelta := lib.ArrIntFind(deltas, 0)

		isSafe := isSameDeltaSign && RedNoseFactorIsSafe(maxDelta) && !containsZeroDelta
		if isSafe {
			numOfSafeReports++
		}

		fmt.Printf(
			"#%d levels: %v, deltas %v (absDelta %v, max: %d) isSafe: %t\n",
			i,
			levels,
			deltas,
			deltasAbs,
			maxDelta,
			isSafe,
		)
	}

	fmt.Printf("Numer of safe reports: %d\n", numOfSafeReports)
}

func processInputPartTwo(input string) {

}
