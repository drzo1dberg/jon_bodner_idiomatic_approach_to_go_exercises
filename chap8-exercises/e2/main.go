package main

import "fmt"

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

type Printable interface {
	fmt.Stringer
	Number
}

func Double[T Number](x T) T {
	return x + x
}

type MyInt int

func (m MyInt) String() string {
	return fmt.Sprintf("%d", m)
}

type MyFloat float64

func (m MyFloat) String() string {
	return fmt.Sprintf("%f", m)
}

func PrintIt[T Printable](v T) {
	fmt.Println(v)
}

func main() {
	PrintIt(Double(MyInt(10)))
	PrintIt(Double(MyFloat(1.87)))
}
