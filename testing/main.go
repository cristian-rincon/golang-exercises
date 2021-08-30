package main

import "fmt"

func Sum(a, b int) int {
	/* Sum two given values */
	return a + b
}

func GetMax(x, y int) int {
	/* Get the max value from two given values */
	if x > y {
		return x
	}
	return y
}

func Fibonacci(n int) int {
	if n <= 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main()  {
	numbers := [10]int{1,2,3,4,5,6,7,8,9,10}
	for _, number := range numbers {
		fmt.Println(Fibonacci(number))
	}
	
}