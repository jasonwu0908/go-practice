package main

import "fmt"

// func main() {
// 	// COMPOSITE LITERAL
// 	var x []int
// 	y := []int{4, 3, 2, 7, 8, 9}
// 	z := append(x, 1, 2, 3, 4)
// 	fmt.Println(y)
// 	fmt.Println(z)
// }

// func main() {
// 	// COMPOSITE LITERAL
// 	var x []int
// 	y := []int{4, 3, 2, 7, 8, 9}
// 	z := append(x, 1, 2, 3, 4)
// 	fmt.Println("y_lens: ", len(y))
// 	fmt.Println("y: ", y)

// 	fmt.Println("=====================")

// 	fmt.Println("y: ", y[3:5])
// 	fmt.Println("len: ", len(y[3:5]))
// 	fmt.Println("cap: ", cap(y[3:5]))
// 	fmt.Println("z: ", z)

// 	fmt.Println("=====================")

// 	y = append(y, z...)
// 	fmt.Println("y_append: ", y)

// 	fmt.Println("=====================")

// 	originYlens := len(y) - len(z)
// 	y = append(y[:originYlens], y[:originYlens]...)
// 	fmt.Println("y_append1: ", y)

// 	for _, v := range y {
// 		fmt.Println(v)
// 	}

// 	fmt.Println("=====================")

// 	for i := 0; i < len(y); i++ {
// 		fmt.Println(y[i])
// 	}
// }

// func main() {
// 	x := make([]int, 10, 12)
// 	fmt.Println(x)
// 	fmt.Println(len(x))
// 	fmt.Println(cap(x))

// 	x = append(x, 12321)
// 	fmt.Println(x)
// 	fmt.Println(len(x))
// 	fmt.Println(cap(x))

// 	x = append(x, 12321)
// 	fmt.Println(x)
// 	fmt.Println(len(x))
// 	fmt.Println(cap(x))

// 	x = append(x, 12321)
// 	fmt.Println(x)
// 	fmt.Println(len(x))
// 	fmt.Println(cap(x))
// }

func main() {
	animal := []string{"dog", "cat", "bird", "mouse"}
	fruit := []string{"apple", "orange", "watermelon", "grape"}
	test := [][]string{animal, fruit}
	fmt.Println(test)
}
