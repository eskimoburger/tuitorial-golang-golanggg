package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func main() {

	pos := addTest()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i))
	}

	//fmt.Println(addVals(1, 2, 3, 5, 6))

}

func addVals(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func addTest() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
