package main

import "fmt"

func main() {
	x := foo()
	fmt.Printf("x type: %T\n", x)
	fmt.Printf("x() type: %T\n", x())
	i := x()
	fmt.Println(i)
}

func foo() func() int {
	return func() int {
		return 666
	}
}
