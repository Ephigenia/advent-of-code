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
	// ids containing repeated patterns are invalid
	pattern, _ := StrContainsRepeatedPattern(str)
	if pattern == "" {
		return true
	}
	return pattern+pattern == str
	// return pattern == "" || (pattern != "" && count == 2)
}

func StrContainsRepeatedPattern(str string) (pattern string, count int) {
	// Check for repeated patterns (11, 1010, 446446, etc.)
	for patternLen := 1; patternLen <= len(str)/2; patternLen++ {
		pattern := str[:patternLen]
		occurrences := len(str) / patternLen
		if (occurrences) > 2 {
			continue
		}
		if strings.Repeat(pattern, occurrences) == str[:len(pattern)*(occurrences)] {
			ptr := str[:len(pattern)*(occurrences)]
			return ptr, strings.Count(str, ptr)
		}
	}
	return "", 0
}

func InvalidIdsFromRange(start, end int) []int {
	var invalidIds []int
	for i := start; i <= end; i++ {
		if !IsValidId(strconv.Itoa(i)) {
			invalidIds = append(invalidIds, i)
		}
	}
	return invalidIds
}
