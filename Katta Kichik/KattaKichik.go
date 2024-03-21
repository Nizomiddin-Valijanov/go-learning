package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var amal string
	var nums = []int{1, 3, 34, 12, 6, 6, 34, 7, 3, 4, 7}
	fmt.Println("Qilinadigon amalni yozing: ")
	fmt.Fscan(os.Stdin, &amal)

	son := KattaKichik(amal, nums)
	if strings.ToLower(amal) == "katta" {
		fmt.Println("Kattasi: ", son)
	} else if strings.ToLower(amal) == "kichik" {
		fmt.Println("Kichigi: ", son)
	}
}

func KattaKichik(amal string, array []int) int {
	var num = array[0]
	switch amal {
	case "katta":
		for i := 1; i < len(array); i++ {
			if array[i] > num {
				num = array[i]
			}
		}
	case "kichik":
		for i := 1; i < len(array); i++ {
			if array[i] < num {
				num = array[i]
			}
		}
	default:
		fmt.Println("Iltimos amalni to'gri kiriting")
	}

	return num
}
