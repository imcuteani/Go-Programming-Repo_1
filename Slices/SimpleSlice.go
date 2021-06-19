// Program to display the simple Slice in Go programming

package main

import "fmt"

func main() {

	arr1 := [4]string{"This", "is", "the", "array"}

	fmt.Println("The array is: ", arr1)

	mySlice := arr1[1:3]
	fmt.Println("The slice is: ", mySlice)

	fmt.Println("The length of the slice is: ", len(mySlice))
	fmt.Println("The Capacity of the slice is: ", cap(mySlice))
}
