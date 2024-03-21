package misollar1

import "fmt"

func main() {
	var age, name = multipleFuncValues(12, 13, "Иван", "Иванов")
	fmt.Println(age, name)
}

// 1.Верните несколько значений функции.
func multipleFuncValues(x, y int, firstName, lastName string) (int, string) {
	var z int = x + y
	var fullName = firstName + " " + lastName
	return z, fullName
}

// 2.Fizz Buzz agar 1 dan 20 bo'lgan sonlar uchga qoldiqsiz bo'linsa Fizz yoqasa Buzz
// agar 5ga ham qoldiqsiz bo'linsa FizzBuzz

func fizzBuzz() {
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println("Buzz")
		}
	}
}
