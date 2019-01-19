package main

import (
	"bank"
	"fmt"
	"time"
)

func main() {
	go bank.Deposit(200)
	go bank.Deposit(100)
	go bank.Deposit(700)

	time.Sleep(1 * time.Second)
	fmt.Println(bank.Balance())
}
