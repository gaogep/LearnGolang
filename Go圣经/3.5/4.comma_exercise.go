package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(myComma("abcsssio"))
}

func myComma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}

	t := n % 3
	z := n - t
	x := 0
	for {
		buf.WriteString(s[x : x+3])
		x += 3
		if x == z {
			break
		}
		buf.WriteByte(',')
	}

	if t != 0 {
		buf.WriteByte(',')
		buf.WriteString(s[x:n])
	}

	return buf.String()
}
