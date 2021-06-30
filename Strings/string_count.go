// Program to show string count function in Go

package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "The world of Python & Golang is same, go language is popular!! $$"
	str2 := "The money $$ for working with Golang!! @@"

	result1 := strings.Count(str1, "go")
	fmt.Println("The result 1 is: ", result1)

	result2 := strings.Count(str2, "$ey")
	fmt.Println("The result 2 is: ", result2)

	result3 := strings.Count(str2, "i")
	fmt.Println("The result 3 is: ", result3)
}
