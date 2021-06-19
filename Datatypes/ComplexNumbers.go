package main

import "fmt"

func main() {

	var a complex128 = complex(4.1, 3.5)
	var b complex64 = complex(3.7, 6.1)

	fmt.Printf("The type of the variable is %T\n", a)
	fmt.Printf("The value of the variable is %d\n", a)
	fmt.Printf("The type of the variable is %T\n", b)
	fmt.Printf("The value of the variable is %d\n", b)
}
