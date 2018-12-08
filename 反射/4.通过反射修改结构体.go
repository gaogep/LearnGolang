package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func main() {
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()

	// 把s.Type()返回的Type对象复制给typeofT，typeofT也是一个反射。
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		// 迭代s的各个域，注意每个域仍然是反射。
		f := s.Field(i)
		// 提取了每个域的名字
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	// 这里T的域的名字都是大写的（被导出的）
	// 因为一个struct中只有被导出的域才是settable的
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)
}
