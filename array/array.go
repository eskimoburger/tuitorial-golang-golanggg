package main

import "fmt"

func main() {
	var a [3]int
	a[0] = 10
	a[1] = 20
	a[2] = 30

	fmt.Println(a)
	fmt.Println(a[2])
	fmt.Println(len(a))

}
