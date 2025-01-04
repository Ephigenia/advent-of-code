package main

import (
	"errors"
	"fmt"
	"os"
	"path"
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
	filename := path.Join(pwd, args[1])

	rawInput, err := lib.ReadInputFile(filename)
	if err != nil {
		lib.ExitWithError(err)
	}
	input := parseInput(rawInput)
	fmt.Printf("Input: %v\n", input)
	printMap(input)
	// processInputPartOne(input)
}

func parseInput(inputs string) []int {
	all := []int{}
	for i := 0; i < len(inputs); i++ {
		c, _ := strconv.Atoi(string(inputs[i]))
		all = append(all, c)
	}
	return all
}

type File struct {
	id     int
	blocks int
	free   int
}

type Disk struct {
	files []File
}

func (d *Disk) addFile(f File) {
	d.files = append(d.files, f)
}

func printMap(m []int) {
	disk := Disk{}

	for i := 0; i < len(m); i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Printf("i: %d\n", i)
		free := 0
		if i < len(m)-1 {
			free = m[i+1]
		}
		disk.addFile(File{
			id:     m[i] / 2,
			blocks: m[i],
			free:   free,
		})
	}

	spew.Dump(disk)

	mode := 1
	parts := []string{}
	id := 0
	for i := 0; i < len(m); i++ {
		part := ""
		length := m[i]
		if mode == 0 { // free space
			part = strings.Repeat(".", length)
		} else { // blocked space
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

// func processInputPartOne(inputs []int) {
// 	spew.Dump(inputs)
// 	printMap(inputs)
// }
