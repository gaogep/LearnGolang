package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5, 6}
	ch := make(chan int)
	for _, x := range list {
		go func(value int) {
			fmt.Println(value)
			ch <- value
		}(x)
	}

	// 等待所有go程结束
	for range list {
		<-ch
	}
}
