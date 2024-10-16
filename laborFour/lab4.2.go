package main

import "fmt"

func average(people map[string]int) float64 {
	var sum int
	var count int
	for _, age := range people {
		sum += age
		count++
	}
	if count == 0 {
		return 0
	}
	return float64(sum) / float64(count)
}
func main() {
	people := map[string]int{
		"vovan": 21, "Kurkin": 23, "jaffar": 27, "haftar": 75,
	}
	avg := average(people)
	fmt.Println("Результат", avg)
}
