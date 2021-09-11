package main

import (
	"fmt"
	"github.com/cristian-rincon/gocalculatorcli/calculator"
)

func main() {
	fmt.Println("Hello, World!")
	app := calculator.App{}
	app.Run()
}