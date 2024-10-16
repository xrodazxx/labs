package main

import "fmt"

func main() {
	var n int
	fmt.Println("количество чисел:")
	fmt.Scanln(&n)
	numbers := make([]int, n)
	fmt.Println("Введите числа:")
	for i := 0; i < n; i++ {
		fmt.Scanln(&numbers[i])
	}
	fmt.Println("Обратный:")
	for i := n - 1; i >= 0; i-- {
		fmt.Println("числа", numbers[i])
	}
}
