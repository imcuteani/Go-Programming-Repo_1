// Program to show simple String Operations

package main

import (
	"fmt"
	"strings"
)

func main() {

	//Comparison between two strings
	strCom := strings.Compare("Adrian", "Andrei")
	fmt.Println("The compared string result is: ", strCom)

	strContain := strings.Contains("The Great City of San Antonio", "Anthony")
	fmt.Println("The substring contain result is: ", strContain)

}
