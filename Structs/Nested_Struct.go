// Program to show nested structure in Go

package main

import "fmt"

type Employee struct {
	Id          int
	FirstName   string
	LastName    string
	Designation string
	salary      string
	JoiningDate string
}

// Nested Struct

type Manager struct {
	Project string
	Band    string
	details Employee
}

func main() {

	NestedStruct := Manager{
		Project: "Data Science",
		Band:    "7",
		details: Employee{303234, "Alan", "Graham", "Project Lead", "$6500", "07/07/2017"},
	}

	fmt.Println("The Name of the Employee is: ", NestedStruct.details.FirstName)
	fmt.Println("The designation of the Employee is: ", NestedStruct.details.Designation)
	fmt.Println("The Project assigned to the employee is: ", NestedStruct.Project)
	fmt.Println("The Band of the employee is: ", NestedStruct.Band)
	fmt.Println("The salary of the employee is: ", NestedStruct.details.salary)

}
