package main

import "fmt"

func main() {
	var n []int
	fmt.Println(n, len(n), cap(n))
	n = append(n, 10)
	fmt.Println(n, len(n), cap(n))
	n = append(n, 20)
	fmt.Println(n, len(n), cap(n))
	n = append(n, 30)
	fmt.Println(n, len(n), cap(n))
	n = append(n, 40)
	fmt.Println(n, len(n), cap(n))
	n = append(n, 50)
	fmt.Println(n, len(n), cap(n))

}
