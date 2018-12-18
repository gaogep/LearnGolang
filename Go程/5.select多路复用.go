package main

import (
	"fmt"
	"time"
)

// 利用select语句重写( 3.和4.) 中的fanin函数
func fanIn_new(input1, input2 chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func SoBoring2(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Second * 2)
		}
	}()
	return c
}

func main() {
	c := SoBoring2("Joe")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		// time.After 在过一段时间之后返回一个通道
		case <-time.After(1 * time.Second):
			fmt.Println("You are too slow")
			return
		}
	}
}

/*  退出select
	-----------------
	quit := meke(chan bool)
	c := boring("Joe", quit)
	for i := rand.Intn(10); i > 0; i-- { fmt.Println(<-c) }
    quit <- true
	-----------------

	In boring function
	------------------
	select {
	case c <- fmt.Sprintf("%s %d", msg, i):
	case <-quit:
		return
	}
	------------------
*/
