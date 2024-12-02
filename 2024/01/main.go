package main

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	pwd, _ := os.Getwd()
	spew.Dump(pwd)
	dat, err := os.ReadFile(pwd + "/input.txt")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(dat))
}
