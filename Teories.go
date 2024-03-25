package main

import (
	"fmt"
	"math"
)

func main() {
	// O'zgaruvchilar
	var a int = 5              // butun son
	const pi float64 = 3.14159 // o'zgarmas qiymat

	// Consolega chiqarish
	fmt.Println("Salom Olam")    // Salom olam
	fmt.Println("Salom \n Olam") // Salom
	// Olam
	fmt.Println("nizomi" + " is cool boy") // nizomi is cool boy
	fmt.Printf("%.2f \n", 2.4434532334)    // 2.44

	// Data typelar
	var b bool = true             // mantiqiy qiymat
	var str string = "nizomiddin" // matn
	var c int = 4                 // butun son
	var d byte = 65               // 8 bitlik butun son
	var e rune = 'A'              // unicode belgilar
	var f float32 = 3.14          // qoldiqli son
	var g complex64 = 3 + 4i      // kompleks son

	fmt.Println(d, e, f, g)

	// Metodlar
	lenStr := len(str)            // uzunlikni aniqlash
	sqrt := math.Sqrt(float64(c)) // ildizdan chiqarish

	fmt.Println(lenStr) // 10
	fmt.Println(sqrt)   // 2

	// Mantiqiy operatorlar
	result := a > c && b == true   // && va operatori
	result2 := a < c || b == false // || yoki operatori
	result3 := a != c              // != teng emas operatori

	fmt.Println(result)  // false
	fmt.Println(result2) // true
	fmt.Println(result3) // true

	// Formatlar
	fmt.Printf("%d\n", a)    // 5
	fmt.Printf("%s\n", str)  // nizomiddin
	fmt.Printf("%.2f\n", pi) // 3.14
	fmt.Printf("%t\n", b)    // true

	// & va * Belgilar
	fmt.Println(&a)   // o'zgaruvchining manzili
	var ptr *int = &a // manzil olish
	fmt.Println(*ptr) // manzil orqali qiymatni olish
}
