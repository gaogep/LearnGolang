package main

import "fmt"

// 这个函数就把right的输出和left的输入联系起来了
func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	const n = 10000
	leftmost := make(chan int)
	right := make(chan int)
	left := leftmost

	// 创建长度为n的daisy链
	// left = leftmost <- right / right <- right1 / right1 <- right2 / right2 <- right3...
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}

	// 在链的最右端输入1，那么最左端就会得到10001
	go func(c chan int) { c <- 1 }(right)
	fmt.Println(<-leftmost)
}
