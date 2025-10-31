package main

import "fmt"

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func Double[T Number](x T) T {
	return x + x
}

type MyInt int

func main() {
	fmt.Println(Double(MyInt(10)))
	fmt.Println(Double(7))
	fmt.Println(Double(1.87))
}
