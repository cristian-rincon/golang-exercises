package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {
}

func parser(input string) int {
	input_cast, _ := strconv.Atoi(input)
	return input_cast
}

func readInput() string {
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	return reader.Text()
}

func (calc) calculate(text string, operator string) int {
	//
	cleanInput := strings.Split(text, operator)
	first := parser(cleanInput[0])
	second := parser(cleanInput[1])

	switch operator {
	case "+":
		fmt.Println(first + second)
		return first + second
	case "-":
		fmt.Println(first - second)
		return first - second
	case "*":
		fmt.Println(first * second)
		return first * second
	case "/":
		fmt.Println(first / second)
		return first / second
	default:
		fmt.Println("Invalid operator")
		return 0
	}

}

func main() {
	user_input := readInput()
	user_operator := readInput()
	c := calc{}
	c.calculate(user_input, user_operator)

}
