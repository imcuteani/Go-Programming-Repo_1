package main

import "fmt"

func main() {
	var language string = "Golang"

	switch {
	case language == "Python":
		{
			fmt.Println("The language selected is Python")
		}

	case language == "Java/J2EE":
		{
			fmt.Println("The language selected is Java/J2EE")
		}

	case language == "Golang":
		{
			fmt.Println("The language selected is Golang!")
		}
	default:
		{
			fmt.Println("The language selected is invalid")
		}
	}
}
