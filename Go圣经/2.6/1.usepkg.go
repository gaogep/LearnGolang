package main

// 在GOPATH中建立了一个包并导入使用
import "tempconv"
import "fmt"

func main() {
	fmt.Println(tempconv.AbsoluteZeroC)
}
