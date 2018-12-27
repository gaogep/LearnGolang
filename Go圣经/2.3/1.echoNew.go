package main

import (
	"flag"
	"fmt"
	"strings"
)

// 创建一个新的布尔标识变量 返回值为指针
var n = flag.Bool("n", false, "omit trailing newline")

// 创建一个新的字符串标识变量 返回值为指针
var sep = flag.String("s", " ", "separator")

func main() {
	// 更新标识变量的默认值 顺带从os.Args[1:]获取命令行参数
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
