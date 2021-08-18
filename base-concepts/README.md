## Repaso de Go

```

package main

import (
	"fmt"
	"strconv"
	"time"
)

func main ()  {
	// Create a new integer [Explicitly]
	// var x int
	// Assign a value to it
	// x = 5
	// Create a new integer [Implicitly]
	// y := 1
	// print the value of x & y
	// fmt.Println(x, y)

	// Error handling
	myVal, err := strconv.ParseInt("2", 0, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(myVal)
	}

	// Working with map
	m := make(map[string]int)
	m["a"] = 1
	fmt.Println(m["a"])

	// Working with slice
	s := []int{1, 2, 3}
	for i, v := range s {
		fmt.Println(i, v)
	}
	// Add a new element to the slice
	s = append(s, 4)
	fmt.Println(s)

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
```