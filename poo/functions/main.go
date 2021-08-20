// Anonymous Functions
package main

import (
	"fmt"
	"time"
)

func main() {
	// x := 6
	// y := func() int {
	// 	return x * 2
	// }()
	// fmt.Println(y)
	c := make(chan int)
	go func() {
		fmt.Println("Starting process")
		time.Sleep(time.Second * 5)
		fmt.Println("Ending process")
		c <- 1
	}()
	<-c
}
