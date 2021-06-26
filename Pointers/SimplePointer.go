//Program to show the Simple Pointer operation in Golang

package main

import "fmt"

func main() {

	var a int = 100

	//Pointer declaration & address by reference

	var b *int = &a

	fmt.Println("The value is the variable b is: ", b)
	fmt.Printf("The data type of the variable b is: %T\n", b)

	// Pointer declaration & address by reference

	var c *int = &a

	fmt.Println("The value of the variable C is: ", c)
	fmt.Printf("The data type of the variable c is: %T\n", c)

}
