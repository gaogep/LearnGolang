package main

import (
	"fmt"
	"math"
)

// 接口Shaper
type Shape interface {
	Area() float32
}

// 类型Square
type Squar struct {
	side float32
}

// 类型Circle
type Circle struct {
	radius float32
}

// Square的Area()方法
func (sq Squar) Area() float32 {
	return sq.side * sq.side
}

// *Circle的Area()方法
func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

// 通过类型断言来判断接口类型sh中的动态类型到底是Squar还是*Circle
func main() {
	sq1 := Squar{5}
	var sh Shape = sq1
	if t, ok := sh.(Squar); ok {
		fmt.Println(t, ok)
	}
	if T, ok := sh.(*Circle); ok {
		fmt.Println(T, ok)
	}

	fmt.Println()
	fmt.Println("-----上方类型断言--------分割线--------下方type_switch-----")
	fmt.Println()

	// 接口变量的类型也可以使用一种特殊形式的 swtich 来检测: type-swtich
	// 如果不赋值 仅仅检测类型的话可以使用-> switch sh.(type)
	switch x := sh.(type) {
	case Squar:
		fmt.Println(x, "is Squar")
	case *Circle:
		fmt.Println(x, "is *Circle")
	}
}

// 此函数用空接口来接收不同的值 再用switch_type来判断类型
func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("Param #%d is a bool\n", i)
		case float64:
			fmt.Printf("Param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("Param #%d is a int\n", i)
		case nil:
			fmt.Printf("Param #%d is a nil\n", i)
		case string:
			fmt.Printf("Param #%d is a string\n", i)
		default:
			fmt.Printf("Param #%d is unknown\n", i)
		}
	}
}
