package main

import "fmt"

func lengthString(a string) int {
	return len(a)

}
func main() {
	var input string
	fmt.Println("Введите строку:")
	fmt.Scan(&input)
	lenght := lengthString(input)
	fmt.Println(lenght)
}
