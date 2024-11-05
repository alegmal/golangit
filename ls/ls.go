package main

import (
	"flag"
	"fmt"
	"os"
)

func checkForError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	showHidden := flag.Bool("a", false, "show hidden files and directories")
	flag.Parse()

	defer fmt.Print("\n")

	files, err := os.ReadDir(".")
	checkForError(err)

	if *showHidden {
		fmt.Print(". .. ")
	}

	for _, file := range files {
		if *showHidden || file.Name()[0] != '.' {
			fmt.Print(file.Name() + " ")
		}
	}
}
