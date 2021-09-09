package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// start := time.Now()
	servers := []string{"http://platzi.com", "http://google.com", "http://udemy.com"}
	channel := make(chan string)

	// for _, server := range servers {
	// 	go checkServer(server, channel)
	// }

	// for i := 0; i < len(servers); i++ {
	// 	fmt.Println(<-channel)
	// }
	i := 0
	for {
		if i>3 {
			break
		}
		for _, server := range servers {
			go checkServer(server, channel)
		}

		time.Sleep(time.Second * 1)
		fmt.Println(<-channel)
		i++
	}

	// end := time.Since(start)
	// fmt.Println("Time elapsed: ", end)
}

func checkServer(server string, channel chan string) {
	// Check if server is up
	_, err := http.Get(server)
	if err != nil {
		channel <- server + " | is down ºnº"
	} else {

		channel <- server + " | is up ºwº"
	}

}
