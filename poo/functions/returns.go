package main

import "fmt"

// Variatic functions
func sum(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

// Returns with names
func getValues(x int) (double int, quad int) {
	double = x * 2
	quad = x * 4
	return
}

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(getValues(3))
}
