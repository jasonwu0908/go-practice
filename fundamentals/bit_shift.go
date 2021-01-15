package main

import (
	"fmt"
	"math"
)

// const (
// 	kb uint32 = 1024
// 	mb uint32 = 1024 * kb
// 	gb uint32 = 1024 * mb
// )

// iota => 1*10, 2*10, 3*10
const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

var x = 2 << 4

func main() {
	gby := float64(gb)
	gbyLog := math.Logb(gby)
	mby := float64(mb)
	mbyLog := math.Logb(mby)
	kby := float64(kb)
	kbyLog := math.Logb(kby)
	fmt.Printf("x: %d\t\t%b\n", x, x)

	fmt.Printf("%d\t\t%b\t", kb, kb)
	fmt.Println("2 ^", kbyLog)
	fmt.Printf("%d\t\t%b\t", mb, mb)
	fmt.Println("2 ^", mbyLog)
	fmt.Printf("%d\t\t%b\t", gb, gb)
	fmt.Println("2 ^", gbyLog)
}
