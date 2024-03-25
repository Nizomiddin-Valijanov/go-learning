package main

import "fmt"

func main() {
	noname := "NNN"
	fmt.Println(noname)
	update(&noname)
	fmt.Println(noname)

}

func pointer() int {
	a := 4
	// pointer := &a
	return *&a
}

func update(str *string) {
	*str = "LLL"
}
