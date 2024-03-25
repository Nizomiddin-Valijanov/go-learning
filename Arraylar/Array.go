package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
}

func ikkiOlchovliArray() {
	array := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 121}}
	fmt.Println(array)
	fmt.Println(array[1][2])
	fmt.Println(array[2][3])
}

func callArray() {
	var array [4]int
	array[3] = 2
	array[3] = 3
	array[3] = 4
}

func OrtaArifmetik() {
	nizoms := [8]float64{1, 4, 3, 6, 4, 5, 7, 6}
	var sum float64 = 0

	for i := 0; i < len(nizoms); i++ {
		sum += nizoms[i]
	}

	var result float64 = sum / float64(len(nizoms))
	fmt.Println(math.Round(result))
}

func Slice() {
	// Slice E'lon qilish
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 8, 9}
	slice = append(slice, 10)
	slice[0] = 777
	sort.Ints(slice)
	// String sort qilish
	strings := []string{"olma", "banan", "ananas", "kiwi", "nok"}
	sort.Strings(strings)
	fmt.Println(strings)
}

func IndexAndItem() {
	slice := []int{1, 2, 9, 5, 6, 7, 8, 9, 0, 7, 6, 3, 5, 7}
	for index, item := range slice {
		fmt.Printf("%d: %d\n", index, item)
	}
}
