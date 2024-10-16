package main

import "fmt"

func main() {
	var people = map[string]int{
		"fork": 32, "forg": 21, "notre": 83,
	}
	delete(people, "forg")
	fmt.Println(people)
}
