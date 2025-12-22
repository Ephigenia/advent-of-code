package main

import (
	"errors"
	"fmt"
	"os"
	"path"

	lib2024 "github.com/Ephigenia/advent-of-code/2024/lib"
	lib2025 "github.com/Ephigenia/advent-of-code/2025/lib"
)

func main() {
	pwd, _ := os.Getwd()
	args := os.Args

	if len(args) < 2 {
		lib2024.ExitWithError(errors.New("please provide a filename"))
	}
	inputFilename := path.Join(pwd, args[1])

	rawInput, err := lib2024.ReadInputFile(inputFilename)
	if err != nil {
		lib2024.ExitWithError(err)
	}

	processInputPartOne(rawInput)
}

// correct answer is 17179
func processInputPartOne(input string) {
	grid := lib2025.NewGridFromString(input)
	grid.Print()

	found := 0
	gridIterator := func(x, y int, value rune) {
		vals := grid.GetAround(x, y)
		count := 0
		for _, v := range vals {
			if v == '@' {
				count++
			}
		}
		if count < 4 {
			found++
		}
	}
	grid.Iterate(gridIterator)

	fmt.Println("Part One:", found)
}
