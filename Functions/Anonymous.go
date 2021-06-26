// Program to show the Anonymous Function

package main

import "fmt"

func main() {

	// Assigning the anonymous function to a variable

	value := func() {
		fmt.Println("Welcome! to Golang World")
	}

	value()

}
