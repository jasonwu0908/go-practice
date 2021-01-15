package main

import "fmt"

func foo() {
	fmt.Println("foo ran")
}

func main() {
	foo()
	func() {
		fmt.Println("Anonymous func ran")
	}()

	func(x int) {
		fmt.Println("Your age:", x)
	}(25)
}
