package main

import (
	"fmt"
	"math/rand"
	"time"
)

func msgGen() chan string {
	out := make(chan string)
	//i := 0
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			out <- fmt.Sprintf("message %d", i)
			i++
		}
	}()
	return out
}
func main() {
	m1 := msgGen()
	m2 := msgGen()
	//m := fanIN(m1, m2)
	mSel := fanBySel(m1, m2)
	for {
		fmt.Println(<-mSel)
	}
}

func fanBySel(m1 chan string, m2 chan string) chan string {
	out := make(chan string)
	go func() {
		for {
			select {
			case n := <-m1:
				out <- n
			case n := <-m2:
				out <- n
			}
		}
	}()
	return out
}

func fanIN(m1 chan string, m2 chan string) chan string {
	out := make(chan string)
	go func() {
		for {
			out <- <-m1
		}
	}()
	go func() {
		for {
			out <- <-m2
		}
	}()
	return out
}
