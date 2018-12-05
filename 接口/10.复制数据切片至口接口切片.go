package main

import "fmt"

// 将数据切片中的数据复制至空接口切片必须要使用range循环
// 否则编译会出错

func main() {
	interfaceSlice := make([]interface{}, 5)
	a := []int{1, 2, 3, 4, 5}
	for i, v := range a {
		interfaceSlice[i] = v
	}
	fmt.Println(interfaceSlice)
}
