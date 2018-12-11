package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	CopyFile("0.noodle.txt", "1.noodle.txt")
	fmt.Println("Copy done")
}

// 利用io包进行文件拷贝
// 注意defer的使用 当打开目标文件的时候发生了错误
// defer最终都会将文件关闭 而不会让文件占用资源
func CopyFile(dstName, srcName string) (wrriten int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
