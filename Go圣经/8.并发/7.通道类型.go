package main

import "fmt"

// out是一个只能发送的通道 不能接收
// 换句话说这是一个只能入的通道
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}

	close(out)
}

// in是一个只能接收的通道 不能发送
// 即一个只能出的通道 不能入
func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}

	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	natu := make(chan int)
	squa := make(chan int)

	go counter(natu)
	go squarer(squa, natu)
	printer(squa)
}
