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
	fmt.Printf("Raw input %s\n", rawInput)

	disk := parseInput(rawInput)
	fmt.Printf("%s\n", disk.String())
}

func parseInput(inputs string) Disk {
	all := lib.ArrStrToInt(strings.Split(inputs, ""))
	disk := Disk{}

	for i := 0; i < len(all); i++ {
		if i%2 == 1 {
			continue
		}
		free := 0
		if i < len(all)-1 {
			free = all[i+1]
		}
		disk.addFile(File{
			id:     all[i] / 2,
			blocks: all[i],
			free:   free,
		})
	}
	return disk
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

func (d *Disk) getFiles() []File {
	return d.files
}

func (d *Disk) String() string {
	str := ""
	for _, file := range d.getFiles() {
		str += strings.Repeat(strconv.Itoa(file.id), file.blocks)
		str += strings.Repeat(".", file.free)
	}
	return str
}
