package main

import "fmt"

func average(a, b, c float64) float64 {
	return (a + b + c) / 3
}
func main() {
	var a, b, c float64
	fmt.Println("Введите три числа:")
	fmt.Scan(&a, &b, &c)
	avg := average(a, b, c)
	fmt.Println("Среднее значение: %.2f\n", avg)
}
