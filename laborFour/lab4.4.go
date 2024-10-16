package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Println("Введите строку")
	fmt.Scanln(&input)
	fmt.Println("Верх регистор", strings.ToUpper(input))
}
