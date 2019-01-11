package main

import "fmt"

type point struct {
	x, y int
}

type colorp struct {
	*point
	color string
}

func (p point) sump(z *point) int {
	return p.x + p.y + z.x
}

func (p colorp) showcolor() {
	fmt.Println(p.color)
}

func main() {
	tp := colorp{&point{1, 2}, "black"}
	tp2 := colorp{&point{1, 2}, "black"}
	p1 := point{1, 2}
	p2 := point{1, 2}
	// 将方法绑定到一个变量上 -> 称之为方法变量
	// 并进行调用
	sum := tp.sump
	// 注意调用的时候参数为*point 不能直接传入colorp
	fmt.Println(sum(tp2.point))

	// 方法表达式 把原来方法的接收者变成了函数的第一个形参
	sumn := point.sump
	fmt.Println(sumn(p1, &p2))
}
