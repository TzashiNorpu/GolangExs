package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}
func worker(id int, c chan int) {
	/*for {
		fmt.Printf("worker %d received %c\n", id, <-c)
	}*/
	for n := range c {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("worker %d received %d\n", id, n)
	}
}
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)
	var values []int //int slice
	n := 0
	tm := time.After(10 * time.Second) // 程序开始运行10秒种后返回一个channel
	tick := time.Tick(2 * time.Second) //每秒钟往通道写一个值
	fmt.Println("Now:", time.Now())
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n = <-c1:
			values = append(values, n)
		case n = <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond): //c1和c2中生成数据的时间超过了800毫秒就输出timeout??
			fmt.Println("timeout")
		case <-tick: //拿到通道值后输出剩余数据的长度
			fmt.Println("Now:", time.Now())
			fmt.Println("queue length:", len(values))
		case <-tm: //这个channel中给有数据时退出程序，也就是10秒中后退出
			fmt.Println("queue left length:", len(values))
			fmt.Println("bye")
			return
		}
	}

}
