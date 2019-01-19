package main

import "sync"

var (
	mu       sync.Mutex
	balbance int
)

func Depo(a int) {
	mu.Lock()
	defer mu.Unlock()
	balbance = balbance + a
}

func Bal() int {
	mu.Lock()
	defer mu.Unlock()
	b := balbance
	return b
}
