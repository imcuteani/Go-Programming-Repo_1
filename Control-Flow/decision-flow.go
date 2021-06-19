package main

import "fmt"

func main() {

	var i int = 100

	if i <= 20 && i >= 40 {
		fmt.Println("The value is i less than actual")
	} else if i >= 40 && i <= 60 {
		fmt.Println("The value is i intermediate")
	} else if i <= 60 && i >= 80 {
		fmt.Println("The value is i is between 60 & 80")
	} else {
		fmt.Println("The value of i is more than 90")
	}

}
