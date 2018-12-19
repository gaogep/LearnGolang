package main

// 通过go程进行多路复用 Multiplex
import (
	"fmt"
	"math/rand"
	"time"
)

// 通道中的值不能直接从一个通道传到另一个通道
func fanIn(input1, input2, input3, input4 <-chan string) <-chan string {
	c := make(chan string)
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
	go func() {
		for {
			c <- <-input3
		}
	}()
	go func() {
		for {
			c <- <-input4
		}
	}()
	return c
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	c := fanIn(boring("Joe"), boring("Ann"), boring("Jack"), boring("Lisa"))
	for i := 0; i < 20; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You are boring. I'm leaving...")
}
