package main

import (
	"fmt"
	"os"
)

func divAndMod(a, b int) (int, int, error) {
	if b == 0 {
		return a, b, fmt.Errorf("dominator is 0")
	}
	return a / b, a % b, nil

}

func main() {
	a, b, err := divAndMod(10, 0)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Println(a, b)

}
