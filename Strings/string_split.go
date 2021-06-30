// Program to show the String split

package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "Welcome to the online portal!"
	str2 := "This is Golang World!"

	result1 := strings.Split(str1, "online")
	fmt.Println("The result 1 is: ", result1)

	result2 := strings.Split(str2, "Go")
	fmt.Println("The result 2 is: ", result2)

	result3 := strings.SplitAfter(str1, "!")
	fmt.Println("The result3 is: ", result3)

	result4 := strings.SplitAfter(str2, "World!")
	fmt.Println("The result 4 is: ", result4)

}
