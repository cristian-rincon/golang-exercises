package main

import "fmt"

func Worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		fib := Fibonacci(j)
		fmt.Println("worker", id, "finished job", j, "result", fib)
		results <- fib
	}
}

func Fibonacci(n int) int {
	if n <= 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 40}
	numWorkers := 4
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for i := 0; i < numWorkers; i++ {
		go Worker(i, jobs, results)
	}

	for _, task := range tasks {
		jobs <- task
	}

	close(jobs)

	for i := 0; i < len(tasks); i++ {
		// fmt.Println(<-results)
		<-results
	}
}
