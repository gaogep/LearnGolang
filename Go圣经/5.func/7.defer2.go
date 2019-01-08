package main

import "fmt"

// return最先执行->return负责将结果写入返回值中->
// 接着defer开始执行一些收尾工作->最后函数携带当前返回值退出

// defer写的时候放在return语句前面
// 执行的时候在return语句之后执行 并可以更新函数的结果变量
// 通过命名结果变量和增加defer语句 可以在每次函数调用的时候输出它的参数和结果
func d(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d", x, result) }()
	return x + x
}

func main() {
	d(5)
}
