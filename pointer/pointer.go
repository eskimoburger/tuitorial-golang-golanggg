package main

import "fmt"

func main() {
	var x int32 = 1
	var y bool = true
	pointerX := &x
	pointerY := &y

	var pointerZ *string
	s := "student"
	pointerZ = &s

	fmt.Println(x, y)
	fmt.Println(*pointerX, *pointerY)
	fmt.Println(&x, &y)
	fmt.Println(pointerX, pointerY)
	fmt.Println(&pointerX, &pointerY)

	fmt.Println(*pointerZ)
	fmt.Println(&pointerZ)

}
