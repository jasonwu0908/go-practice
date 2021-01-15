package main

import "fmt"

type person struct {
	first  string
	last   string
	gender string
	age    int
	phone  string
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	s1 := secretAgent{
		person: person{
			first:  "Jason",
			last:   "Wu",
			gender: "m",
			age:    25,
			phone:  "0926000001",
		},
		ltk: false,
	}
	fmt.Println(s1)
	fmt.Println(s1.person)
	fmt.Println(s1.person.first, s1.first)

}
