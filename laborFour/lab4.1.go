package main

import "fmt"

func main() {
	people := map[string]int{"Gnom": 192, "Kuror": 27, "Corp": 29}
	people["Horse"] = 15
	fmt.Println(people)
	for name, age := range people {
		fmt.Println(":  лет\n", name, age)
	}

}
