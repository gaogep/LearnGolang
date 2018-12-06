package main

import (
	"fmt"
	"reflect"
)

// 反射可以在运行时检查类型和变量，例如它的大小、方法和 动态 的调用这些方法。
// 反射可以从接口值反射到对象，也可以从对象反射回接口值
func main() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("value:", v.Float())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)
}
