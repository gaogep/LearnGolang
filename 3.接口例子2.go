package main

import "fmt"

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price
}

type valuable interface {
	getValue() float32
}

func showValue(asset valuable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}

// 函数showValue的参数为一个asset的接口类型 而stockPositon和car都实现了getValue()
// 方法，所以可以被赋值给接口类型o 从而调用各自的getValue()方法
// 可以整么说: 所有实现了valuable接口的类型都可以调用showValue()这个函数
func main() {
	var o valuable = stockPosition{"GOOG", 577.20, 4}
	showValue(o)
	o = car{"BMW", "M3", 66500}
	showValue(o)
}
