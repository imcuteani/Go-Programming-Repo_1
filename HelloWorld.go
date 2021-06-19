//Declaration of the package
package main

import "fmt"

// Declaration of the main() method

func main() {

	// Declaration of the variables

	var a int = 100
	var b float64 = 66.39
	var lang string = "Golang"
	var c int = 401

	// Declaration of the values of the variables

	fmt.Printf("Hi! This is your first Go Program\n")
	fmt.Printf("Hi! This is to print your Hello World Program\n")
	fmt.Printf("Lets write some numbers %d\n", a)
	fmt.Printf("Lets write some precision numbers %f\n", b)
	fmt.Printf("Lets write some string %s\n", lang)
	fmt.Printf("Lets write anything %v\n", c)

	// Lets print the type of the variables

	fmt.Printf("The type of the 1st variable is %T\n", a)
	fmt.Printf("The type of the 2nd variable is %T\n", b)
	fmt.Printf("The type of the 3rd variable is %T\n", c)
	fmt.Printf("The type of the 4th variable is %T\n", lang)
	fmt.Println("This is the end of the Hello World Program")
}
