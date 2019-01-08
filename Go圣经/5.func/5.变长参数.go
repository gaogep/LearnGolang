package main

import "fmt"

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))

	// 当实参存在于一个slice中的时候 调用的时候在后面加省略号
	list := []int{1, 2, 3, 4}
	fmt.Println(sum(list...))
}
