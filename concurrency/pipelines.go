package main

import "fmt"

func Generator(c chan<- int) {
	for i := 0; i <= 10; i++ {
		c <- i
	}
	close(c)
}

func double(in <-chan int, out chan<- int) {
	for v := range in {
		out <- v * 2
	}
	close(out)
}

func Print(c chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func main() {
	generator := make(chan int)
	doubleChan := make(chan int)
	go Generator(generator)
	go double(generator, doubleChan)
	Print(doubleChan)
}
