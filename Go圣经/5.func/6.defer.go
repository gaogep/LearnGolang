package main

import (
	"log"
	"time"
)

func SlowOperation() {
	// 延迟执行的是trace调用之后返回的匿名函数
	// trace并不会延迟调用 所以会先打印enter的消息
	defer trace("SlowOperation")()
	time.Sleep(time.Second * 10)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	// 延迟执行的函数在return语句之后执行
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

func main() {
	SlowOperation()
}
