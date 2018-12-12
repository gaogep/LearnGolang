package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 不适用缓冲
	fmt.Fprintf(os.Stdout, "%s\n", "hello world - unbuffered")

	// os.Stdout实现了Write方法
	buf := bufio.NewWriter(os.Stdout)

	fmt.Fprintf(buf, "%s\n", "hello world - buffered")
	// 向缓冲区写入以后刷新缓冲区结果才会显示出来
	buf.Flush()
}
