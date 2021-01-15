package main

import "fmt"

// func main() {
// 	switch {
// 	case (1 == 2):
// 		fmt.Println("this should not print")

// 	case (2 == 3):
// 		fmt.Println("this should not print")

// 	case (3 == 3):
// 		fmt.Println("this is 3 == 3")
// 		fallthrough

// 	case (4 == 4):
// 		fmt.Println("this is 4 == 4")
// 		fallthrough

// 	case (5 == 6):
// 		fmt.Println("this is 5 == 6")
// 		fallthrough

// 	case (6 == 7):
// 		fmt.Println("this is 6 == 7")
// 		fallthrough

// 	case (7 == 7):
// 		fmt.Println("this is 7 == 7")

// 	default:
// 		fmt.Println("this is default !")
// 	}
// }

func main() {
	n := "Test1"
	switch n {
	case "Jason":
		fmt.Printf("I'm %v\n", n)
	case "Joe":
		fmt.Printf("I'm %v\n", n)
	case "Test":
		fmt.Printf("I'm %v\n", n)
	default:
		fmt.Printf("%v not in case\n", n)
	}
}
