package main

import (
	"fmt"
	"io"
	"net/http"
)

type WebWriter struct {}

func (w WebWriter) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	r, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	e := WebWriter{}
	io.Copy(e, r.Body)
	

}
