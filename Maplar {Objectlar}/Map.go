package main

import "fmt"

func main() {
	var maps = createMap()
	UpdateMap(maps)
	deleteMap(maps)
	el, status := maps["dollars"]
	fmt.Println(maps)
	fmt.Println(el, status)
}

func createMap() map[string]int {
	var money map[string]int = map[string]int{
		"dollars": 1233,
		"euros":   2000,
		"apples":  3000,
	}
	return money
}

func UpdateMap(maps map[string]int) map[string]int {
	maps["dollars"] = 7700
	return maps
}

func deleteMap(maps map[string]int) map[string]int {
	delete(maps, "apples")
	return maps
}
