package main

import (
	"fmt"
)

// func main()  {
// 	for i := 0; i<=100; i++{
// 		fmt.Println(i)
// 	}
// }

// func main()  {
// 	x := 1
// 	for x < 10 {
// 		fmt.Println(x)
// 		x++
// 	}
// 	fmt.Println("done.")
// }

// func main()  {
// 	x := 1
// 	for {
// 		if x > 9 {
// 			break
// 		}
// 		fmt.Println(x)
// 		x++
// 	}
// 	fmt.Println("done.")
// }

func main() {
	x := 0
	for {
		x++
		if x > 100 {
			break
		}

		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
	fmt.Println("done!")
}
