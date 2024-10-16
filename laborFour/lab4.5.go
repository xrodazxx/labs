package main

import "fmt"

func main() {
	var n, sum int
	fmt.Println("Введите несколько чисел (0 для конца):")
	for {
		fmt.Println("Введите число:")
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		sum += n
	}
	fmt.Println("Сумма введённых чисел:", sum)
}
