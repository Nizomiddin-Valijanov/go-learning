package main

import "fmt"

func main() {
	var nums = []any{1, 2, true, 2, 54, "Salom", 2, 4, 4.53, 6, "assdf", 45, 3, false, 3, 34.34, "Nizomi", 34, "olam", 8}
	sorting(nums)
}

func sorting(array []any) {
	var (
		nums    []int
		strings []string
		bools   []bool
		floats  []float64
	)

	for _, item := range array {
		switch item.(type) {
		case int:
			nums = append(nums, item.(int))
		case string:
			strings = append(strings, item.(string))
		case bool:
			bools = append(bools, item.(bool))
		case float64:
			floats = append(floats, item.(float64))
		}
	}

	fmt.Println("Целые числа:", nums)
	fmt.Println("Строки:", strings)
	fmt.Println("Bools:", bools)
	fmt.Println("Floats:", floats)
}
