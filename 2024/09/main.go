package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"
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
	input := parseInput(rawInput)
	fmt.Printf("Input: %v\n", input)
	printMap(input)
	processInputPartOne(rawInput)
}

func printMap(m []int) {
	mode := 1
	parts := []string{}
	id := 0
	for i := 0; i < len(m); i++ {
		part := ""
		length := m[i]
		if mode == 0 {
			part = strings.Repeat(".", length)
		} else {
			part = strings.Repeat(strconv.Itoa(id), length)
		}
		parts = append(parts, part)
		if i%2 == 0 {
			mode = 0
		} else {
			id++
			mode = 1
		}
	}
	fmt.Printf("%s\n", strings.Join(parts, ""))
}

func parseInput(inputs string) []int {
	all := []int{}
	for i := 0; i < len(inputs); i++ {
		c, _ := strconv.Atoi(string(inputs[i]))
		all = append(all, c)
		// fmt.Printf("%d\n", c)
	}
	return all
}

func processInputPartOne(inputs string) {

}
