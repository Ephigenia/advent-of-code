package main

import (
	"errors"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"

	lib2024 "github.com/Ephigenia/advent-of-code/2024/lib"
	"github.com/Ephigenia/advent-of-code/2025/lib"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	pwd, _ := os.Getwd()
	args := os.Args

	if len(args) < 2 {
		lib.ExitWithError(errors.New("please provide a filename"))
	}
	inputFilename := path.Join(pwd, args[1])

	rawInput, err := lib.ReadInputFile(inputFilename)
	if err != nil {
		lib.ExitWithError(err)
	}

	processInputPartOne(rawInput)
}

// right answer is 13919717792
func processInputPartOne(input string) {
	re := regexp.MustCompile("\r?\n")
	normalizedInput := re.ReplaceAllString(input, "")
	parts := strings.Split(normalizedInput, ",")
	invalidIds := []int{}
	for _, part := range parts {
		splitted := lib2024.ArrStrToInt(strings.Split(part, "-"))
		invalidIds = append(invalidIds, InvalidIdsFromRange(splitted[0], splitted[1])...)
	}
	spew.Dump(invalidIds)
	sum := lib2024.ArrIntSum(invalidIds)
	spew.Dump("sum", sum)
}

func IsValidId(str string) bool {
	// ids starting with zero are invalid
	if str[0:1] == "0" {
		return true
	}
	size := len(str) / 2
	chunks := []string{
		str[0:size],
		str[size:],
	}
	if chunks[0] == chunks[1] {
		return false
	}
	return true
}

func InvalidIdsFromRange(start, end int) []int {
	invalidIds := make([]int, 0)
	for i := start; i <= end; i++ {
		if !IsValidId(strconv.Itoa(i)) {
			invalidIds = append(invalidIds, i)
		}
	}
	return invalidIds
}
