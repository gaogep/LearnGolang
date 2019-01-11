package main

import "math"

type Point struct {
	x, y float64
}

// 参数p称为方法的接收者
// 表达式p.Distance称作选择子 因为它为接收者p选择合适的Distance方法
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

// 用指针来传递变量的地址用于更新接收者 避免复制整个实参
// 下面的方法名字为(*Point).ScaleBy 括号是必须的 没有括号表达式
// 会被解析为*(Point.ScaleBy)
func (p *Point) ScaleBy(factor float64) {
	p.x *= factor
	p.y *= factor
}

// 注意 不允许本身是指针的类型进行方法声明
// 不能对一个取地址的Point接收者参数调用*Point方法 因为无法获取临时变量的地址
// -> Point{1, 2}.ScaleBy(2) 编译错误！！！ 不能获得Point字面量的地址

// 编译器会自动添加& * 从而达到T类型变量 *T类型变量二者相互调用对方的方法！！！

func main() {

}
