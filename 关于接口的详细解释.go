package main

import "io"

// 从概念上来讲，一个接口类型的值简称接口值，其实有两个部分组成：
// 一个具体类型和该类型的一个值。二者称为接口的动态类型和动态值

// 声明了一个接口类型的值w，w的动态类型和值都为nil 一个接口值是否为nil
// 取决于它的动态类型，所以w现在是一个nil接口值 可以用 == != 来检测一
// 个接口的值是否为nil，调用一个nil接口的任何方法 都会导致崩溃
var w io.Writer

// 现在 把一个*os.File类型的值赋值给w 这次赋值把一个具体类型隐式转换
// 为一个接口类型，它与对应的io.Writer(os.stdout)相同
// 此外 不管这种转换是隐式的还是显式的 它都可以转换操作数的类型和值
// 接口值的动态类型会设置为指针类型*os.File的类型描述符，它的动态值
// 会设置为os.Stdout的副本，即一个指向代表进程标准输出的os.File指针

w = os.Stdout


// 调用该接口值的Write方法实际会调用(*os.File).Write的方法 即输出 "hello"
w.Write([]byte("hello"))

// 上面的就等价于
os.Stdout.Write([]byte("hello"))
