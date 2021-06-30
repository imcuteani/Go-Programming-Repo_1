//package to show golang program to string trimming

package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "I love Golang programming!"
	str2 := "$$Golang was the most adorable language of 2020!!"

	result1 := strings.Trim(str1, "!")
	fmt.Println("The trimmed string 1 is: ", result1)

	result2 := strings.Trim(str2, "$!")
	fmt.Println("The trimmed string 2 is: ", result2)

	result3 := strings.TrimLeft(str1, "!")
	fmt.Println("The result 3 is: ", result3)

	result4 := strings.TrimLeft(str2, "$$")
	fmt.Println("The trim left function is: ", result4)

}
