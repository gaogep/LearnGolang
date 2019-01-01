package main

import "fmt"

// nonempty 返回一个新的slice slice中的元素都是非空字符串
// 在函数调用的过程中 底层数组发生了改变
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}

	return strings[:i]
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Println(data)
	data = nonempty(data)
	fmt.Println(data)
}
