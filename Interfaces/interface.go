package main

import "fmt"

type Numbers struct {
	a int
	b int
}

type Calculating interface {
	Sum() int
	Multiply() int
	Division() float64
	Substract() int
}

func (n Numbers) Sum() int {
	return n.a + n.b
}
func (n Numbers) Multiply() int {
	return n.a * n.b
}

func (n Numbers) Division() float64 {
	return float64(n.a) / float64(n.b)
}
func (n Numbers) Substract() int {
	return n.a - n.b
}

// Consolas, 'Courier New', monospace
func main() {
	var i Calculating
	number := Numbers{1, 2}
	i = number
	fmt.Println("plus:", i.Sum())
	fmt.Println("minus:", i.Substract())
	fmt.Println("ko'paytirish:", i.Multiply())
	fmt.Println("bo'lish:", i.Division())
}
