package main

import "fmt"

// 一个接口的值可以赋值给另一个接口变量，只要底层类型实现了必要的方法。
type SqrInterface interface {
	Sqr() float32
}
type AbsInterface interface {
	Abs() int
}
type Point struct{ X, Y int }

var ai AbsInterface
var si SqrInterface

func main() {
	pp := new(Point)
	var empty interface{}

	// 空接口可以接收任何类型的值
	empty = pp

	ai = empty.(AbsInterface) // 如果pp实现了Abs()赋值会成功否则失败
	si = empty.(SqrInterface) // 同理对SqrInterface也是一样的

	empty = si

	fmt.Println(ai)
}
