package main

import "fmt"

func main() {
	foo()
	bar("bar")
	x, y := woo("Jason", "Ian")
	fmt.Println(x, y)
}

func foo() {
	fmt.Println("foo!")
}

func bar(s string) {
	fmt.Println("Hello, ", s)
}

func woo(x string, y string) (string, bool) {
	a := fmt.Sprint(x, " ", y, `, says "Hello"`)
	b := true
	return a, b
}
