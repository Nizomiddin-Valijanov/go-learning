package main

import (
	"fmt"
)

func main() {
	module("salom")
	var arr [10]int
	arr[0] = 15
	arr[1] = 25
	arr[2] = 35
	arr[3] = 45

	// var nums = []int{1, 2, 2, 3, 42, 2, 2, 5, 7, 5, 6, 8, 5, 8, 6, 8}
	var hello = "salom olam "
	var str_arr = []string{hello}
	fmt.Println(str_arr, hello[0])
	var age, name = multipleFuncValues(12, 13, "Иван", "Иванов")
	fmt.Println(age)
	fmt.Println(name)
}

// 1.Верните несколько значений функции.
func multipleFuncValues(x, y int, firstName, lastName string) (int, string) {
	var z int = x + y
	var fullName = firstName + " " + lastName
	return z, fullName
}

// 2.Инициализируйте структуру.

// func mostFrequent(array []int) {
// 	frequency := make(map[int]int)
// 	for _, num := range array {
// 		frequency[num]++
// 	}

// 	maxFrequency := 0

// 	for _, freq := range frequency {
// 		if freq > maxFrequency {
// 			maxFrequency = freq
// 		}
// 	}

// 	for num, freq := range frequency {
// 		if freq == maxFrequency {
// 			fmt.Println("The most-repeated:", num)
// 		}
// 	}

// 	fmt.Println("Count of repeat:", maxFrequency)
// }

// func fizzBuzz() {
// 	for i := 1; i <= 20; i++ {
// 		if i%3 == 0 && i%5 == 0 {
// 			fmt.Println("FizzBuzz")
// 		} else if i%3 == 0 {
// 			fmt.Println("Fizz")
// 		} else {
// 			fmt.Println("Buzz")
// 		}
// 	}
// }

// func sorting(array []any) {
// 	var (
// 		nums    []int
// 		strings []string
// 		bools   []bool
// 		floats  []float64
// 	)

// 	for _, item := range array {
// 		switch item.(type) {
// 		case int:
// 			nums = append(nums, item.(int))
// 		case string:
// 			strings = append(strings, item.(string))
// 		case bool:
// 			bools = append(bools, item.(bool))
// 		case float64:
// 			floats = append(floats, item.(float64))
// 		}

// 	}

// 	fmt.Println("Целые числа:", nums)
// 	fmt.Println("Строки:", strings)
// 	fmt.Println("Bools:", bools)
// 	fmt.Println("Floats:", floats)
// }

// func KattaKichik(amal string, array []int) int {
// 	var num = array[0]
// 	switch amal {
// 	case "katta":
// for i := 1; i < len(array); i++ {
// 	if array[i] > num {
// 		num = array[i]
// 	}
// }
// 	case "kichik":
// 		for i := 1; i < len(array); i++ {
// if array[i] < num {
// 	num = array[i]
// }
// 		}
// 	default:
// 		fmt.Println("Iltimos amalni to'gri kiriting")
// 	}

// 	fmt.Println(num)
// 	return num
// }
