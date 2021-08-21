package main

import "testing"

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		cc              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			cc: "1",
			mockFunc: func() {
				GetEmployeeByID = func(id int) (Employee, error) {
					return Employee{
						Id:       1,
						Position: "CEO",
					}, nil
				}

				GetPersonByCC = func(id string) (Person, error) {
					return Person{
						CC:  "1",
						Name: "John Doe",
						Age:  25,
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					CC:  "1",
					Name: "John Doe",
					Age:  25,
				},
				Employee: Employee{
					Id:       1,
					Position: "CEO",
				},
			},
		},
	}
	originalGetEmployeeByID := GetEmployeeByID
	originalGetPersonByCC := GetPersonByCC
	for _, test := range table {
		test.mockFunc()
		ft, err := GetFullTimeEmployee(test.id, test.cc)

		if err != nil {
			t.Errorf("Error when getting Employee")
		}

		if ft.Age != test.expectedEmployee.Age {
			t.Errorf("Error, got %d expected %d", ft.Age, test.expectedEmployee.Age)
		}
	}

	GetEmployeeByID = originalGetEmployeeByID
	GetPersonByCC = originalGetPersonByCC

}
