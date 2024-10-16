package main

import "fmt"

func differAndSum(a, b float64) (float64, float64) {
	dif := a - b
	sum := a + b
	return sum, dif
}
func main() {
	a := 34.3
	b := 13.56
	sum, dif := differAndSum(a, b)
	fmt.Printf("Сумма: %.2f\n", sum)
	fmt.Printf("Разность: %.2f\n", dif)
}
