package main

import (
	"fmt"
	"io"
	"os"
)

// 类型断言是一个作用在接口值上的操作，即x.(T)，其中x是一个接口类型的表达式
// 而T是一个类型，称为断言类型。类型断言会检查作为操作数的动态类型是否满足指
// 定的断言类型

// 为了判断一个接口值是否保存了一个特定的类型，类型断言可返回两个值:
// 其底层值以及一个报告断言是否成功的布尔值

// 判断接口值w是否保存了*os.File类型
func main() {
	var w io.Writer
	w = os.Stdout
	f, ok := w.(*os.File)
	fmt.Println(f, ok)
	// 成功 f == Stdout
	// output->&{0xc04206a1e0} true
}
