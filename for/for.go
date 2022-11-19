package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])

	// }
	// i := 0
	// for i < len(numbers) {
	// 	fmt.Println(numbers[i])
	// 	i++
	// }
	// for i, v := range numbers {
	// 	fmt.Println(i, v)
	// }

	i := 0
	for {
		fmt.Println(numbers[i])
		i++
		if i == len(numbers) {
			break
		}
		fmt.Println(i)
	}

}
