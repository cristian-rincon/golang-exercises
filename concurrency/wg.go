package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup // Declaramos una variable de tipo WaitGroup, esto nos permitirá bloquear el programa hasta que terminen las rutinas
		
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomething(i, &wg)
	}

	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	duration := 2 * time.Second
	fmt.Printf("Iteración %v\n", i)
	fmt.Println("Comenzando proceso de 2 segundos")
	fmt.Printf("Comenzando proceso de %v\n segundos.", duration)
	time.Sleep(duration)
}
