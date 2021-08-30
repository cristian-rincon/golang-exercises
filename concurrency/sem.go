package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan int, 5) // Buffer de 5
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		c <- 1
		wg.Add(1)
		go doSomething2(i, &wg, c)
	}
	wg.Wait()
}

func doSomething2(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	duration := 4 * time.Second
	fmt.Printf("Iteración %v inicializada\n", i)
	time.Sleep(duration)
	fmt.Printf("Iteración %v finalizada\n", i)
	<-c
}
