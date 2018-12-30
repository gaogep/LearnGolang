package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := "abc"

	// []byte(s)转换操作会分配新的字节数组，拷贝添入
	// s含有的字节，并生成一个slice引用， 指向整个数组
	// 具备优化功能的编译器在某些情况下可能会避免分配内存和复制内容
	b := []byte(s)

	fmt.Println(s, b)
	fmt.Println(intsToString([]int{1, 2, 3}))
}

// bytes包和strings包都预备了许多对应的实用函数 它们两两对应
// bytes包和stings包都具备了以下6个函数
// Contains Count Fields HasPrefix Index Join
// 唯一不同的是，操作对象由字符串变成了字节slice

// bytes包为高效处理字节slice提供了Buffer类型
// Buffer起始为空，其大小随着各种类型数据的写入而增长
// Buffer也无需初始化，原因是0值本来就有效
// 例如以下这个函数 与 Sprintf类似，不过加上了逗号
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
