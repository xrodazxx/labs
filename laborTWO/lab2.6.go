package main

import "fmt"

func avg(a, b int) int {
	return (a + b) / 2
}
func main() {
	var a, b int
	fmt.Println("Введите 2 числа")
	fmt.Scan(&a, &b)
	average := avg(a, b)
	fmt.Println("Средне", average)
}
