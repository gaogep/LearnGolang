package main

import (
	"fmt"
	"strings"
)

func main() {
	t := []string{"abc", "efg", "hik"}
	fmt.Println(strings.Join(t, "-"))
}
