package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 第三个参数 读文件的时候被忽略 写文件的时候无论是Windows还是Linux都是0666
	OpFile, OpErr := os.OpenFile("noodle2.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if OpErr != nil {
		fmt.Println(OpErr)
		return
	}

	defer OpFile.Close()
	OpWriter := bufio.NewWriter(OpFile)
	OpString := "hello world!\n"

	// fmt.Fprintf(OpFile, "Some test data.\n")
	OpWriter.WriteString(OpString)
	OpWriter.Flush()

	// 也可以不使用缓冲区 直接向文件中写入
	f, _ := os.OpenFile("test", os.O_CREATE|os.O_WRONLY, 0666)
	defer f.Close()
	f.WriteString("hello, world in a file\n")
}
