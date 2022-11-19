package main

import "fmt"

func main() {
	// var (
	// 	a int64
	// 	b int64
	// )
	// a := 10
	// b := 2
	// c := a + b

	// fmt.Printf("%d\n", a+b)
	// fmt.Printf("%d\n", a-b)
	// fmt.Printf("%d\n", a*b)
	// fmt.Printf("%d\n", a/b)
	// fmt.Printf("%d\n", a%b)

	a := 0b1111
	b := 0b0101

	fmt.Printf("%04b\n", a&b)
	fmt.Printf("%04b\n", a|b)
	fmt.Printf("%04b\n", a^b)

}
