package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	files, err := os.ReadDir(".")
	check(err)

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}
}
