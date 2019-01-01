package main

import "fmt"

type Point struct{ x, y int }

func main() {
	// 建议使用以下这种方式来初始化一个结构体变量
	// 即通过指定部分或者全部成员变量的名称来对结构体进行初始化
	p := Point{x: 1, y: 2}
	t := fmt.Sprintf("%q", p)
	fmt.Println(t)
}
