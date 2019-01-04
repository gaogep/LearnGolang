package main

import "fmt"

// 一个函数如果又命名的返回值，可以省略return语句的操作数，这称之为裸返回
func nake(a, b int) (c int) {
	c = a + b
	return
}

func main() {
	fmt.Println(nake(2, 3))
}
