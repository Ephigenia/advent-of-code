package main

import (
	"errors"
	"os"
	"path"
	"regexp"
	"strconv"
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
	for _, part := range parts {
		splitted := strings.Split(part, "-")
		spew.Dump(splitted)
	}
}

func IsValidId(id int) bool {
	strId := strconv.Itoa(id)
	if strId[0:1] == "0" {
		return false
	}
	return !StrContainsRepeatedPattern(strId)
}

func StrContainsRepeatedPattern(str string) bool {
	// Check for repeated patterns (11, 1010, 446446, etc.)
	for patternLen := 1; patternLen <= len(str)/2; patternLen++ {
		pattern := str[:patternLen]
		if strings.Repeat(pattern, len(str)/patternLen) == str[:len(pattern)*(len(str)/patternLen)] {
			return true
		}
	}
	return false
}

func InvalidIdsFromRange(start, end int) []int {
	var invalidIds []int
	for i := start; i <= end; i++ {
		if StrContainsRepeatedPattern(strconv.Itoa(i)) {
			invalidIds = append(invalidIds, i)
		}
	}
	return invalidIds
}
