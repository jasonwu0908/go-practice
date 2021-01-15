package main

import "fmt"

func main() {
	// x := [5]int
	var x [5]int
	fmt.Println(x)
	x[3] = 100
	fmt.Println(x)
	fmt.Println(len(x))
}
