package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5}
	sum := foo("Jason", xi...)
	fmt.Println(sum)
}

func foo(name string, x ...int) int {
	fmt.Println(x, name)
	fmt.Printf("%T\n", x)

	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
