package main

import "fmt"

const (
	a = iota
	b
	c
)

const (
	d = iota
	e
	f
)
const (
	// i = 1 << iota
	j = 3 << iota
	k
	l
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	// fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
}
