package main

import "fmt"

type BasicPoint struct {
	x, y int
}

type Circle struct {
	BasicPoint
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	a := Wheel{Circle{BasicPoint{2, 3}, 4}, 20}

	b := Wheel{
		Circle: Circle{
			BasicPoint: BasicPoint{x: 4, y: 7},
			Radius:     3,
		},
		Spokes: 5, // 注意:尾部的逗号是必须的
	}
	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", b) // # 使得printf的格式化符号%v以类似
	// Go语法的方式输出对象，这个方式里面包含了成员变量的名字
}
