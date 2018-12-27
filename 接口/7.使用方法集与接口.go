package main

import "fmt"

// List类型
type List []int

// List的方法
func (L List) Len() int        { return len(L) }
func (L *List) Append(val int) { *L = append(*L, val) }

// 接口Appender 需要实现方法Append
type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i := start; i < end; i++ {
		a.Append(i)
	}
}

// 接口Lener需要实现方法Len()
type Lener interface {
	Len() int
}

func LongEnough(L Lener) bool {
	return L.Len()*10 > 42
}

func main() {

	var lst List

	// CountInto(lst, 1, 10)
	// 在lst上调用CountInto函数会出现编译错误 因为CountInto需要一个Appender
	// 而lst的Append方法定义在指针上
	// 而调用LongEnough不会报错是因为Len方法定义在值上
	if LongEnough(lst) {
		fmt.Printf("- lst is long enough\n")
	}

	// new函数返回一个指针指向plst
	// 调用CountInto是可以的调用LongEnough的时候 指针会被自动解引用
	plst := new(List)
	CountInto(plst, 1, 10)
	if LongEnough(plst) {
		fmt.Printf("- plst is long enough\n")
	}
}

// 总结一下
// 1.指针方法可以通过指针调用
// 2.值方法可以通过值调用
// 3.接收者是值的方法可以通过指针调用，因为指针会首先被解引用
// 4.接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址
