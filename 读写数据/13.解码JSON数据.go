package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
	// 不用理解这个数据的结构，我们可以直接使用Unmarshal 把这个数据编码并保存在接口值中
	var f interface{}
	if err := json.Unmarshal(b, &f); err != nil {
		fmt.Println(err)
	}

	// f 指向的值是一个map，key 是一个字符串，value 是自身存储作为空接口类型的值：
	// 要访问这个数据，我们可以使用类型断言 m变成了映射类型
	m := f.(map[string]interface{})
	fmt.Println(m) // -> map[Name:Wednesday Age:6 Parents:[Gomez Morticia]]
}
