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

	defer File.Close()
	FileReader := bufio.NewReader(File)

	// 一次读取1024字节
	buff := make([]byte, 1024)

	for {
		// ReadNum代表读取到的字节数
		ReadNum, RdErr := FileReader.Read(buff)
		if RdErr == io.EOF && ReadNum == 0 {
			break
		}
	}

	fmt.Println(string(buff))
}
