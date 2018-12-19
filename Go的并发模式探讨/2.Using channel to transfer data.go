package main

// 利用go程实现生产者与消费者
import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go Producer(ch)
	go Consumer(ch)
	time.Sleep(5)
}

func Producer(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func Consumer(ch chan int) {
	for {
		input := <-ch
		fmt.Println(input)
	}
}
