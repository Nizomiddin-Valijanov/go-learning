package main

import "fmt"

type User struct {
	name     string
	age      int64
	password string
}

func main() {
	// var user User = User{name: "John", age: 23, password: "salom"}
	user := User{"John", 23, "qwerty"}
	updateUser(&user)
	fmt.Println(user)
}

func updateUser(user *User) {
	user.name = "Geniy"
	user.age = 234
}
