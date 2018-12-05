package main

import (
	"fmt"
	"os"
)

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

func main() {
	fmt.Println("Please enter your full name: ")

	if _, err := fmt.Scanln(&firstName, &lastName); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName)

	if _, err := fmt.Sscanf(input, format, &f, &i, &s); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("From the string we read: ", f, i, s)
	// 输出结果: From the string we read: 56.12 5212 Go
}
