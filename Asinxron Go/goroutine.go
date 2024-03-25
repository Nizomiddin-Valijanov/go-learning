package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	ch2 := make(chan int)
	time.Sleep(4 * time.Second)
	go say("Salom", ch, ch2)
	fmt.Println(<-ch, <-ch2)
}

func say(msg string, ch chan string, ch2 chan int) {
	fmt.Println(msg)
	ch <- "golangdan salomlar"
	ch2 <- 29
}
