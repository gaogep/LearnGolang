package main

import (
	"fmt"
	"strings"
)

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot > 0 {
		fmt.Println(dot)
		s = s[:dot]
	}
	return s
}

func main() {
	fmt.Println(basename("c.d.go"))
}
