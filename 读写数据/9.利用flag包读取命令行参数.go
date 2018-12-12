package main

import (
	"flag"
	"os"
)

// 定义了一个默认值是false的flag: 当在
// 命令行出现了第一个参数（这里是"n"），flag 被设置成true
// 要给flag 定义其它类型，可以使用flag.Int() ，flag.Float64() ，flag.String()
var NewLine = flag.Bool("n", false, "print newline")

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	flag.PrintDefaults()
	// 扫描参数列表
	flag.Parse()
	var s string = ""

	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "

			// flag 被解引用到
			// *NewLine ，所以当值是true 时将添加一个Newline（"\n"）
			if *NewLine {
				s += Newline
			}
		}

		// flag.Arg(0) 就是第一个真实的flag，而不是像os.Args(0) 放置程序的名字。
		s += flag.Arg(i)
	}

	os.Stdout.WriteString(s)
}
