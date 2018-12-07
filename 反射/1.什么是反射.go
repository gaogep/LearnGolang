package main

import (
	"fmt"
	"reflect"
)

// 反射可以在运行时检查类型和变量，例如它的大小、方法和动态的调用这些方法。
// 反射可以从接口值反射到对象，也可以从对象反射回接口值
func main() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())

	// 类型的分类为Kind->基础类型 数字类型 聚合类型 引用类型 接口类型等
	fmt.Println("kind:", v.Kind())
	fmt.Println("value:", v.Float())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)
}
