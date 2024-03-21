package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math_ceil())
}

func math_1() int {
	a := 10
	b := 14
	return a + b
}

func math_sqrt() float64 {
	var a float64 = 144
	result := math.Sqrt(a)
	return result
}

func math_ceil() int {
	var a float64 = 14.344
	result := math.Ceil(a)
	return int(result)
}

func math_floor() int {
	var a float64 = 14.344
	result := math.Floor(a)
	return int(result)
}

func math_round() int {
	var a float64 = 12.5
	result := math.Round(a)
	return int(result)
}

func Calc() {
	var (
		action string
		first  float64
		second float64
	)
	var result float64
	fmt.Println("Kalkulator")
	fmt.Println("Qaysi amalni bajarmoqchisiz? (+, -, *, /)")
	fmt.Scan(&action)
	fmt.Println("Birinchi son:")
	fmt.Scan(&first)
	fmt.Println("Ikkichi son:")
	fmt.Scan(&second)
	switch action {
	case "+":
		result = first + second
	}
}
