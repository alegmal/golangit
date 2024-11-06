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
	listLong := flag.Bool("l", false, "show more details")
	flag.Parse()

	defer fmt.Print("\n")

	files, err := os.ReadDir(".")
	checkForError(err)

	if *showHidden {
		fmt.Print(". .. ")
	}

	for _, file := range files {
		if *showHidden || file.Name()[0] != '.' {
			if *listLong {
				info, err := os.Stat(file.Name())
				checkForError(err)

				fmt.Println(info.Mode(), " ", info.Size(), info.ModTime().Month(), file.Name(), " ")
			} else {
				fmt.Print(file.Name(), " ")
			}
		}
	}

}
