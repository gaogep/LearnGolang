package main

import (
	"errors"
	"fmt"
)

func main() {
	// 采用errors包的New方法 返回一个err的类型
	err := errors.New("This is a error")
	// 由于已经实现了error接口的方法 因此可以直接调用对应的方法
	fmt.Println(err.Error())

	// 采用fmt.Errof 将string信息转化为error信息 并返回
	err = fmt.Errorf("%s", "the error test for fmt.Errorf")
	fmt.Println(err.Error())
}
