package main

import (
	"fmt"
	"io/ioutil"
)

// 一次性把文件全部读入buf中
func main() {
	File := "D:\\Learn\\Golang\\读写数据\\noodle.txt"
	buf, err := ioutil.ReadFile(File)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))
}
