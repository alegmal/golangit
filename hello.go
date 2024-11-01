package main

import "fmt"

func main() {
	fmt.Println("hello world")

	var potatoe = "potatoe"
	avok := "avocado"

	fmt.Println(potatoe + " " + avok)

	const a = "banana"

	// a = "guyava"

	fmt.Println(a)

	// n := 5
	for n := range 3 {
		fmt.Println(n)
	}

	b := [...]int{100, 11: 400, 500}
    fmt.Println("idx:", b)
}
