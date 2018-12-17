package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

// 指定使用核心的数量
var numCores = flag.Int("n", 3, "number of CPU cores to use")

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(*numCores)

	// Go程并行
	fmt.Println("In main")
	go longwait()
	go shortwait()
	fmt.Println("About to sleep in main")
	time.Sleep(10 * 1e9)
	fmt.Println("At the end of main")
}

func longwait() {
	fmt.Println("Beginning longWait()")
	time.Sleep(5 * 1e9) // 睡眠5s
	fmt.Println("End of longWait()")
}

func shortwait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * 1e9) // 睡眠2s
	fmt.Println("End of shortWait()")
}
