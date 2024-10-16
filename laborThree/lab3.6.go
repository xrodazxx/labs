package main

import (
	"fmt"
)

func main() {
	words := []string{"алынтеев", "муханбедов", "кучеров", "сирентеев", "кучкоглазов"}

	var longest string

	for _, word := range words {
		if len(word) > len(longest) {
			longest = word
		}
	}
	fmt.Println("Самая длинная строка:", longest)
}
