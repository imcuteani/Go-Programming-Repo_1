package main

import "fmt"

// Range function with For loop

func main() {

	rvariable := []string{"My", "name", "is", "Gilbert"}

	// i variable will store the index of the string
	// j variable will store the value of the string

	for i, j := range rvariable {
		fmt.Println(i, j)
	}
}
