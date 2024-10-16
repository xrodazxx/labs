package main

import "fmt"

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Массив")
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}
