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

	// a, b := value()

	// fmt.Println(a,b)

	// sum(1, 2, 3, 4, 5)

	// fmt.Println(fib(7))

	// i := 1
    // fmt.Println("initial:", i)

    // zeroval(i)
    // fmt.Println("zeroval:", i)

    // zeroptr(&i)
    // fmt.Println("zeroptr:", i)

    // fmt.Println("pointer:", &i)

	var buffer [256]byte

	slice := buffer[10:20]
    // for i := 0; i < len(slice); i++ {
    //     slice[i] = byte(i)
    // }
    // fmt.Println("before", slice)
    // AddOneToEachElement(slice)
    // fmt.Println("after", slice)

	// fmt.Println("Before: len(slice) =", len(slice))
    // newSlice := SubtractOneFromLength(slice)
    // fmt.Println("After:  len(slice) =", len(slice))
    // fmt.Println("After:  len(newSlice) =", len(newSlice))

	fmt.Println("Before: len(slice) =", len(slice))
    PtrSubtractOneFromLength(&slice)
    fmt.Println("After:  len(slice) =", len(slice))
}
func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}
func SubtractOneFromLength(slice []byte) []byte {
    slice = slice[0 : len(slice)-1]
    return slice
}
func PtrSubtractOneFromLength(slicePtr *[]byte) {
    slice := *slicePtr
    *slicePtr = slice[0 : len(slice)-1]
}


// func potatoe(a, b, c int) int {
// 	return a+b+c
// }

func sum(nums ...int) {
	fmt.Println(nums)
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func fib(num int) int {
	if num < 2 {
		return num
	}

	return fib(num - 1) + fib(num - 2)
}

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}