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

// any types that have method "speak" also type "human"
type human interface {
	speak()
}

// method
// s		   -> variable
// secretAgent -> type
func (s secretAgent) speak() {
	fmt.Printf(`%v say: "say yolo again!"%v`, s.name, "\n")
}

func (p person) speak() {
	fmt.Printf(`%v say: "say yolo again!"%v`, p.name, "\n")
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into bar ------>", h.(person).name)
	case secretAgent:
		fmt.Println("I was passed into bar ------>", h.(secretAgent).person.name)
	}
	fmt.Println("I was passed into bar", h)
}

func main() {
	sa1 := secretAgent{
		person: person{
			name:   "Jason",
			age:    25,
			gender: "m",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			name:   "Wendy",
			age:    31,
			gender: "w",
		},
		ltk: false,
	}

	p1 := person{
		name:   "Dr. Strange",
		age:    40,
		gender: "m",
	}
	// fmt.Printf("%T\n", p1)
	// fmt.Printf("%T\n", sa1)
	// sa1.speak()
	// sa2.speak()

	bar(sa1)
	bar(sa2)
	bar(p1)
}
