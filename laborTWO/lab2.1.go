package main

import "fmt"

func main() {
	var numb int
	fmt.Println("Введите число:")
	fmt.Scan(&numb)
	if numb%2 == 0 {
		fmt.Println("Число:", numb, " четное")
	} else {
		fmt.Println("Число:", numb, " не четное")
	}
}
