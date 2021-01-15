package main

import "fmt"

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Jason",
		last:  "Wu",
		age:   25,
	}
	fmt.Println(p1)
}
