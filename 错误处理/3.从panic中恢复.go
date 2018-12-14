package main

import "fmt"

// 正如名字一样，这个（recover）内建函数被用于从panic 或错误场景中恢复：
// 让程序可以从panicking 重新获得控制权，停止终止过程进而恢复正常执行。
// ------------------------------------------------------------
// recover 只能在defer 修饰的函数中使用：
// 用于取得panic 调用中传递过来的错误值，如果是正常执行，
// 调用recover 会返回nil，且没有其它效果。
// -------------------------------------------------------------
// 总结：panic 会导致栈被展开直到defer 修饰的recover() 被调用或者程序中止。

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()
	badCall()
	fmt.Printf("After bad call\r\n")
}

func main() {
	fmt.Printf("Calling test\r\n")
	test()
	fmt.Printf("Test completed\r\n")
}
