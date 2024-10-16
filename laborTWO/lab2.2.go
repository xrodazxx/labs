package main

import "fmt"

func checkNumber(number int) string {
	if number > 0 {
		return "Positive"
	} else if number == 0 {
		return "Zero"
	} else {
		return "Negative"
	}
}
func main() {
	var number int
	fmt.Println("Введите значение:")
	fmt.Scan(&number)
	result := checkNumber(number)
	fmt.Println(result)

}
