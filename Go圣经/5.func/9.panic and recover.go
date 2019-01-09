package main

import "fmt"

// 如果内置的recover函数在延迟函数内部调用，且这个包含defer语句的
// 函数发生宕机 recover会终止当前的宕机状态并且返回宕机的值
// 函数不会从之前宕机的地方继续运行 而是正常返回
// 如果recover在其他任何情况下运行则它没有任何效果并且返回nil
func main() {
	d := test()
	fmt.Println(d)
}

func test() (res int) {
	var ok bool
	defer func() {
		if p := recover(); p != nil {
			res, ok = p.(int)
			fmt.Println(ok)
		}
	}()

	panic(5)
}
