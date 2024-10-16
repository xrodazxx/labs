package main

import (
	"fmt"
	"laborThree/mathutils"
	"laborThree/stringutils"
)

func main() {
	var munber int
	var stroke string
	fmt.Println("Введите число:")
	fmt.Scan(&munber)
	fmt.Println(mathutils.Factorial(munber))
	fmt.Println("Введите строку:")
	fmt.Scan(&stroke)
	fmt.Println(stringutils.Revers(stroke))

}
