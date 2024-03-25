package main

import "fmt"

// Function Hoisting Bo'lishini isbotlash
func console(msg any) {
	fmt.Println(msg)
}

func main() {
	handleFunc := func(x int, y int) int {
		return x + y
	}
	square := func(x int) int {
		return x * x
	}
	console(handleFunc(3, 2))
	test(square)
	// a := funcInfunc("Salom")
	// a()
	funcInfunc("salom")()
}

func test(smf func(int) int) {
	fmt.Println(smf(5))
}

// funksiyani ichidagi funksiaya
// funksiya qaytaradi
func funcInfunc(x string) func() {
	return func() {
		fmt.Println(x)
	}
}
