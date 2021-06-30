// Program to show the string comparison using slice

package main

import (
	"fmt"
	"strings"
)

func main() {

	myslice := []string{"Golang", "Python", "Java", "SpringBoot", "netcore"}

	fmt.Println("The initial slice is: ", myslice)

	result1 := "Golang" > "Go"
	fmt.Println("The result1 is: ", result1)

	result2 := "Python" > "Golang"
	fmt.Println("The result2 is: ", result2)

	result3 := "Golang" < "Go"
	fmt.Println("The result3 is: ", result3)

	result4 := "Golang" > "Java"
	fmt.Println("The result4 is: ", result4)

	result5 := "Java" > "SpringBoot"
	fmt.Println("The result5 is: ", result5)

	result6 := "SpringBoot" < "netcore"
	fmt.Println("The result6 is: ", result6)

	// Print by using compare() method

	fmt.Println(strings.Compare("Golang", "Go"))
	fmt.Println(strings.Compare("Java", "Go"))
	fmt.Println(strings.Compare("geeks", "geek"))
	fmt.Println(strings.Compare("springboot", "spring"))
}
