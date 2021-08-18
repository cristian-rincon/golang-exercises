package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

func (ftEmpl FullTimeEmployee) getMsg() string {
	return "Full time employee"
}

type TempEmployee struct {
	Person
	Employee
	taxRate int
}

func (tEmpl TempEmployee) getMsg() string {
	return "Temporary employee"
}

type PrintInfo interface {
	getMsg() string
}

func getMsg(p PrintInfo)  {
	fmt.Printf("%v\n", p.getMsg())
}

func main() {
	ftEmpl := FullTimeEmployee{}
	ftEmpl.name = "John"
	ftEmpl.age = 30
	ftEmpl.id = 1
	fmt.Printf("%v\n", ftEmpl)

	tEmpl := TempEmployee{}
	getMsg(tEmpl)
	getMsg(ftEmpl)
}
