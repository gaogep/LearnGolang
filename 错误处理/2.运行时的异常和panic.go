package main

import (
	"fmt"
	"os"
)

var user = os.Getenv("USER")

// panic 可以直接从代码初始化：当错误条件（我们所测试的代码）很严苛且不可恢复，
// 程序不能继续运行时，可以使用panic 函数产生一个中止程序的运行时错误。
func main() {
	fmt.Println("Starting the program")
	panic("Test panic")
}

// 一个检查程序是否被已知用户启动的具体例子
func check() {
	if user == "" {
		panic("Unknown user: no value for $USER")
	}
}
