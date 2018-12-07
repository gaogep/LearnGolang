package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// Any把任何值格式化为一个字符串
func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

func formatAtom(v reflect.Value) string {
	// 类型的分类为Kind->基础类型 数字类型 聚合类型 引用类型 接口类型等
	switch v.Kind() {
	case reflect.Invalid:
		return "Invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	// case:...后续一些情况先省略
	default:
		return v.Type().String() + " value"
	}
}

func main() {
	x := []int{1, 2, 3}
	fmt.Println(Any(x))
	// -> 结果: []int value
}
