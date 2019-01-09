package main

import (
	"fmt"
)

// 一个典型的宕机发生的时候 所有的延迟函数都会倒序执行
// 然后程序会异常退处并留下一条日志消息
func main() {
	defer fmt.Println("aaa")
	panic("panic")
}
