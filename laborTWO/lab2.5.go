package main

import "fmt"

type rectangle struct {
	bot int
	top int
}

func main() {
	var sqare = rectangle{bot: 8, top: 7}
	fmt.Println("Sqare rectangle:", sqare.bot*sqare.top)
}
