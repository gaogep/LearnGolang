package main

import "fmt"

// 必须通过自定义函数来比较两个map是否相等
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for i, xv := range x {
		// 施用!ok来区分元素不存在 和 元素存在但是为0的情况
		if yv, ok := y[i]; !ok || xv != yv {
			return false
		}
	}
	return true
}

func equal2(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for i, xv := range x {
		fmt.Println(i, xv, y[i])
		if xv != y[i] {
			return false
		}
	}
	return true
}

func main() {
	ages1 := map[string]int{"zpf": 24}
	ages2 := map[string]int{"zpf": 24}
	fmt.Println(equal(ages1, ages2))
	fmt.Println(equal2(map[string]int{"A": 0}, map[string]int{"B": 42}))
	// 由于第二个映射中不存在键为A的值，所以会返回0从而和被equal2判定为2者相等
	// 所以在判断两个map是否相等的时候要注意判断键是否存在 然后再对值进行比较
}
