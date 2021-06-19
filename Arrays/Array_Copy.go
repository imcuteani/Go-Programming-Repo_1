package main

import "fmt"

func main() {

	// Copying an array by using value

	myarr1 := [6]int{30, 40, 50, 60, 70, 80}
	fmt.Println("The first Array is :", myarr1)

	myarr2 := myarr1
	fmt.Println("The second Array is :", myarr2)

	myarr1[4] = 110

	fmt.Println("The updated value of first array is:", myarr1)
	fmt.Println("The updated value of the second array is:", myarr2)

	// Copy array pass by reference

	arr1 := [5]string{"This", "is", "the", "New", "York"}
	fmt.Println("The first array is : ", arr1)
	arr2 := &arr1

	fmt.Println("The copied second array is: ", *arr2)
	arr1[4] = "Jersey"
	fmt.Println("The updated first array is :", arr1)
	fmt.Println("The updated second array is :", *arr2)
}
