package main

import (
	"errors"
	"fmt"
	"os"
	"path"

	lib2024 "github.com/Ephigenia/advent-of-code/2024/lib"
	lib2025 "github.com/Ephigenia/advent-of-code/2025/lib"
	"github.com/davecgh/go-spew/spew"
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

	s := grid.GetS(0, 0)
	spew.Dump(s)
	fmt.Println("Part One:", 0)
}
