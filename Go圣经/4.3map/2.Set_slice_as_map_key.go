package main

import "fmt"

// slice不能用于比较，如何使用slice作为map的key呢？
// 可以定义一个辅助函数，将slice转化为字符串
// 同样的方法可以用于任何不能比较的类型
func transfer(list []string) string { return fmt.Sprintf("%q", list) }
func Add(list []string)             { m[transfer(list)]++ }
func Count(list []string) int       { return m[transfer(list)] }

var m = make(map[string]int)

func main() {
}
