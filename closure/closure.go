package main

import "fmt"

func add() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}

}

func main() {
	pos := add()

	for i := 2; i < 10; i++ {
		fmt.Println(pos(i))
	}

}
