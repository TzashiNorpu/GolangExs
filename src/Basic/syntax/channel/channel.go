package main

import (
	"fmt"
	"time"
)

func worker1(id int, c chan int) {
	for {
		fmt.Printf("worker %d received %c\n", id, <-c)
	}
}
func main() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker1(i, channels[i])
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}
