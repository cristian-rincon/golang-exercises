package main

import "time"

type Person struct {
	CC   string
	Name string
	Age  int
}

type Employee struct {
	Id       int
	Position string
}

type FullTimeEmployee struct {
	Employee
	Person
}

var GetPersonByCC = func(cc string) (Person, error) {
	time.Sleep(5 * time.Second)
	var person Person
	person.CC = cc
	person.Name = "John Doe"
	person.Age = 30
	return person, nil
}

var GetEmployeeByID = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	var employee Employee
	employee.Id = id
	employee.Position = "Full Time"
	return employee, nil
}

func GetFullTimeEmployee(id int, cc string) (FullTimeEmployee, error) {
	var ftemployee FullTimeEmployee
	e, err := GetEmployeeByID(id)
	if err != nil {
		return ftemployee, err
	}
	ftemployee.Employee = e
	p, err := GetPersonByCC(cc)
	if err != nil {
		return ftemployee, err
	}
	ftemployee.Person = p
	return ftemployee, nil
}
