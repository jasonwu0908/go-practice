package main

import "fmt"

type person struct {
	name   string
	age    int
	gender string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Printf(`%v say: "say yolo again!"%v`, s.name, "\n")
}

func main() {
	p1 := secretAgent{
		person: person{
			name:   "Jason",
			age:    25,
			gender: "m",
		},
		ltk: true,
	}
	p2 := secretAgent{
		person: person{
			name:   "Wendy",
			age:    31,
			gender: "w",
		},
		ltk: false,
	}
	p1.speak()
	p2.speak()
}
