// Program to show simple slice operation in Go

package main

import "fmt"

func main() {

	arr := [6]string{"This", "is", "the", "State", "of", "Utah"}

	// The following slice denotes the low bound of arr[1] & upto arr[4]
	var slice_1 = arr[1:5]
	// The following slice denotes the arrays excluding arr[5]
	var slice_2 = arr[:5]
	// The following slice denotes the arrays including arr[1]
	var slice_3 = arr[1:]
	// The following slice denotes the whole slice of array
	var slice_4 = arr[:]

	fmt.Println("The display of simple slice 1: ", slice_1)
	fmt.Println("The display of simple slice 2: ", slice_2)
	fmt.Println("The display of simple slice 3: ", slice_3)
	fmt.Println("The display of simple slice 4: ", slice_4)

}
