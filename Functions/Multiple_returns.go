// Program to show multiple return values in function

package main

import "fmt"

func vals() (int, int) {
	return 4, 6
}

func main() {
	a, b := vals()

	fmt.Println("The value is: ", a)
	fmt.Println("The value is: ", b)

	_, c := vals()
	fmt.Println("The value is: ", c)

}
