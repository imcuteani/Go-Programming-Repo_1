// Program to show if the given characters are present in string

package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "Welcome to Hawaii island!"
	str2 := "Lets learn the Golang programming!"

	fmt.Println("Original strings")
	fmt.Println("String 1:", str1)
	fmt.Println("String 2:", str2)

	// checking if the string present

	result1 := strings.Contains(str1, "Haw")
	result2 := strings.Contains(str2, "Go")

	fmt.Println("The result 1 is: ", result1)
	fmt.Println("The result 2 is: ", result2)

	// Checking with ContainsAny() function

	result3 := strings.ContainsAny(str1, "is")
	result4 := strings.ContainsAny(str2, "prog")

	fmt.Println("The result 3 is: ", result3)
	fmt.Println("The result 4 is: ", result4)
}
