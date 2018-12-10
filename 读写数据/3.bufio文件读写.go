package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	File, Err := os.Open("D:\\Learn\\Golang\\读写数据\\noodle.txt")
	if Err != nil {
		fmt.Println(Err)
		return
	}

	FileReader := bufio.NewReader(File)
	defer File.Close()

	for {
		// 一行行读取文件
		Str, RdErr := FileReader.ReadString('\n')
		fmt.Println(Str)
		if RdErr == io.EOF {
			return
		}
	}
}
