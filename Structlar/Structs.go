package main

import "fmt"

type User struct {
	name     string
	age      int64
	password string
	scores   []int
}

func (u User) getHighScore() int {
	high := 0
	for _, score := range u.scores {
		if score > high {
			high = score
		}
	}
	return high
}

func main() {
	// var user User = User{name: "John", age: 23, password: "salom"}
	user := User{"John", 23, "qwerty", []int{1, 23, 45, 76, 34, 63, 35, 56, 34, 64, 54}}
	updateUser(&user)
	fmt.Println(user)
	fmt.Println(user.getHighScore())
}

func updateUser(user *User) {
	user.name = "Geniy"
	user.age = 234
}
