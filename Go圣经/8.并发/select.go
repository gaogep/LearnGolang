package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	fmt.Println("Press return  to abort")
	tick := time.Tick(1 * time.Second)
	for count := 50; count > 0; count-- {
		fmt.Println(count)
		select {
		case <-tick:
		case <-abort:
			fmt.Println("Lanch aborted")
			return
		}
	}
}
