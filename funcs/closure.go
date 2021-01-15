package main

import "fmt"

func main() {
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(factorial(10))
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func factorial(num int) int {
	if num == 0 {
		return 0
	}
	fmt.Println("factorial:", num)
	return num + factorial(num-1)
}
