package main

import "fmt"

func main() {
}

func loop_1() {
	for i := 0; i <= 100; i++ {
		if i == 50 {
			break
		}
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

func loop_2() {
	var (
		i int = 0
		k int = 0
	)
	for i < 5 {
		i++
		k++
		fmt.Println(k)
	}
}

func matrix() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for _, row := range matrix {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}
		fmt.Println()
	}
}
