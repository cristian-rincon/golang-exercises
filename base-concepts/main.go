package main

import (
	"fmt"
	"time"
)

func main() {

	// Working with channel
	c := make(chan int)
	go doSomeStuff(c)
	<-c

	// Working with pointers
	g := 25
	fmt.Println(g)
	h := &g // h is a pointer to g [Address of g]
	fmt.Println(h)
	fmt.Println(*h) // *h is the value of g [Value of g]
}

func doSomeStuff(c chan int) {
	/*
		Sample code to send a value to the channel
	*/
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1
}
