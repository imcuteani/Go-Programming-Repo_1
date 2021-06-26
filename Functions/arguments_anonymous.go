//Program to Show arguments passing in Anonymous function

package main

import "fmt"

func main() {

	// Passing arguments in anonymous function

	func(element string) {
		fmt.Println(element)
	}("This is Seattle")
}
