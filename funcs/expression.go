package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("my first expression")
	}
	g := func(x int) {
		fmt.Println("your age:", x)
	}
	f()
	g(25)
}
