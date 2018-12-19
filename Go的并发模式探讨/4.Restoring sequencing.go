package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func main() {
	c := FanIn(SoBoring("AAA"), SoBoring("BBB"))
	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		// 当msg1/msg2全部打印完毕以后向 waitForIt 传入 true
		// 使得该通道退处阻塞状态 继续向下执行
		msg1.wait <- true
		msg2.wait <- true
	}
}

func SoBoring(msg string) <-chan Message {
	waitForIt := make(chan bool)
	c := make(chan Message)
	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s %d", msg, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			<-waitForIt // 阻塞直到true传入进来 然后抛弃
		}
	}()
	return c
}

func FanIn(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}
