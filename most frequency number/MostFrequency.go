package main

import "fmt"

func main() {

	var nums = []int{7, 2, 4, 7, 3, 8, 8, 8, 5, 4, 3, 2, 7, 7}
	mostFrequent(nums)
}

func mostFrequent(array []int) {
	frequency := make(map[int]int)
	for _, num := range array {
		frequency[num]++
	}

	maxFrequency := 0

	for _, freq := range frequency {
		if freq > maxFrequency {
			maxFrequency = freq
		}
	}

	for num, freq := range frequency {
		if freq == maxFrequency {
			fmt.Println("The most-repeated:", num)
		}
	}

	fmt.Println("Count of repeat:", maxFrequency)
}
