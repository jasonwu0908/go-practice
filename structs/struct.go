package main

import "fmt"

type person struct {
	first  string
	last   string
	gender string
	age    int
	phone  string
}

func main() {
	p1 := person{
		first:  "Jason",
		last:   "Wu",
		gender: "m",
		age:    25,
		phone:  "0926000001",
	}
	p2 := person{
		first:  "Jow",
		last:   "Wu",
		gender: "m",
		age:    30,
		phone:  "0926111111",
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first)
	fmt.Println(p2.first)
}
