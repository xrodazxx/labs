package main

import "fmt"

func main() {
	users := []string{"ton", "BTC", "BNB", "Hot", "ETH"}
	slice := users[:]
	fmt.Println("Срез:", slice)
	slice = append(slice, "NOT", "TRC")
	fmt.Println("Добавление", slice)
	index := 1
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Println("Delete", slice)

}
