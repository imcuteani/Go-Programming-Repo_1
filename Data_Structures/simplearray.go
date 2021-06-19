package main

import "fmt"

func main() {

	arr := [6]int{3, 5, 7, 9, 11, 13}
	fmt.Printf("The elements of the array is %v\n", arr)
	fmt.Printf("The type of the array variable is %T\n", arr)

	for i := 1; i <= 20; i++ {
		fmt.Println("The elements of the array is ", arr[i])
	}

}
