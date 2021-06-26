// Program to show Simple Struct in Go

package main

import "fmt"

type address struct {
	Name    string
	City    string
	ZipCode string
	Street  string
}

func main() {

	var a address
	fmt.Println("The address is: ", a)

	a1 := address{Name: "Andy", City: "Glasglow", ZipCode: "34421", Street: "Main Street"}

	fmt.Println("The first address is: ", a1)
	fmt.Println("The Name of the person is: ", a1.Name)
	fmt.Println("The City of the person is: ", a1.City)

	a2 := address{Name: "Adrian", City: "London", ZipCode: "32245", Street: "Waterloo Street"}
	fmt.Println("The next address is: ", a2)

	a3 := address{Name: "Andrei", City: "Lisbon"}
	fmt.Println("The last address is: ", a3)

}
