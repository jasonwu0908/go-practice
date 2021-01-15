package main

// ******  1  ******
// func main() {
// 	x := 123
// 	y := "Jason Wu"
// 	z := true
// 	fmt.Println(x, y, z)
// }
// ======== Ans =======
// 123 Jason Wu true
// ====================

// ******  2  ******
// var x int
// var y string
// var z bool

// func main() {
// 	fmt.Println(x)
// 	fmt.Println(y)
// 	fmt.Println(z)
// }
// ======== Ans =======
// 0

// false
// ====================

// ******  3  ******
// var x int = 123
// var y string = "Jason Wu"
// var z bool = true

// func main() {
// 	a := fmt.Sprintln("a:", x, y, z)
// 	b := fmt.Sprintf("b: %v\t%v\t%v", x, y, z)
// 	fmt.Println(a)
// 	fmt.Println(b)
// }
// ======== Ans =======
// a: 123 Jason Wu true

// b: 123  Jason Wu        true
// ====================

// ******  4  ******
// type hotdog int

// var x hotdog
// var y int

// func main() {
// 	fmt.Println(x)
// 	fmt.Printf("%T\n", x)
// 	x := 42
// 	fmt.Println(x)
// 	y = int(x)
// 	fmt.Println(y)
// 	fmt.Printf("%T\n", y)
// 	z := int(x)
// 	fmt.Println(z)
// 	fmt.Printf("%T\n", z)
// }

// ======== Ans =======
// 0
// main.hotdog
// 42
// 42
// int
// 42
// int
// ====================
