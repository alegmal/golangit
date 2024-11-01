package main

import "fmt"

func main() {
	// fmt.Println("hello world")

	// var potatoe = "potatoe"
	// avok := "avocado"

	// fmt.Println(potatoe + " " + avok)

	// const a = "banana"

	// // a = "guyava"

	// fmt.Println(a)

	// // n := 5
	// for n := range 3 {
	// 	fmt.Println(n)
	// }

	// b := [...]int{100, 11: 400, 500}
    // fmt.Println("idx:", b)


	// s := make([]string, 3)
	// fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
	// s[0] = "ok"
	// s[1] = "two"
	// s[2] = "potatoe"
	// // s[3] = "amb"
	// // s[4] = "kkkk"
	// fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
	// fmt.Println(potatoe(1,2,3))

	a, b := value()
	
	fmt.Println(a,b)
}

func potatoe(a, b, c int) int {
	return a+b+c
}
