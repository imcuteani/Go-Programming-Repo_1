package main

import "fmt"

func main() {

	first_str1 := "Golang"
	second_str2 := "Kubernetes"
	third_str3 := "Docker"

	result1 := first_str1 == second_str2
	result2 := second_str2 == third_str3
	result3 := third_str3 == first_str1

	// Display the results

	fmt.Printf("The first & second string comparison result is %s\n", result1)
	fmt.Printf("The type of the first & second string comparison result is %T\n", result1)
	fmt.Printf("The type of the second & third string comparison result is %s\n", result2)
	fmt.Printf("The type of the second & third string result variable is %s\n", result2)
	fmt.Printf("The first & third string comparison result value is %s\n", result3)
	fmt.Printf("The type of the first & third string comparing result variable is %s\n", result3)

}
