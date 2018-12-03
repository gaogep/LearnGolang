package main

// 通过接口来实现多态
import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

// Square的方法接收者为一个指针
func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

// Rectangle的方法接收者为一个Rectangle类型的变量
func (re Rectangle) Area() float32 {
	return re.length * re.width
}

func main() {
	s := &Square{5}
	r := Rectangle{3, 4}
	shapes := []Shaper{s, r}
	for n := range shapes {
		fmt.Println("Shape details:", shapes[n])
		fmt.Println("Area of this shape is:", shapes[n].Area())
	}
}
