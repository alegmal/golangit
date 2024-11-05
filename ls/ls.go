package main

import (
	"fmt"
	"os"
)

func checkForError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	files, err := os.ReadDir(".")
	checkForError(err)

	for _, file := range files {
		// fmt.Println(file.Name(), file.IsDir())
		fmt.Print(file.Name(), " ")
	}
}
