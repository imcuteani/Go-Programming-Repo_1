// Program to show Anonymous Struct in Go

package main

import "fmt"

func main() {

	// Creating Anonymous function

	Student := struct {
		Id        int
		FirstName string
		LastName  string
		language  string
		dept      string
	}{
		Id:        001,
		FirstName: "George",
		LastName:  "Collins",
		language:  "Espanol",
		dept:      "CSE101",
	}

	fmt.Println("The Student details are: ", Student)
	fmt.Println("The Id of the student is: ", Student.Id)
	fmt.Println("The Name of the student is: ", Student.FirstName, " ", Student.LastName)
	fmt.Println("The department of the student is: ", Student.dept)
}
