package main

import "fmt"

func main() {
	// %*s号中的*号输出带有可变数量空格的字符串
	// 宽度和字符串分别由 数字 和 “ ”来控制
	fmt.Printf("%*s + |\n", 1, " ")
	fmt.Printf("%*s + |\n", 2, " ")
	fmt.Printf("%*s + |\n", 3, " ")
	fmt.Printf("%*s + |\n", 4, " ")
	fmt.Printf("%*s + |\n", 5, " ")
}
