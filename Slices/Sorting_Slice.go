// Program to show the sorting of the slice

package main

import (
	"fmt"
	"sort"
)

func main() {

	slice_1 := []string{"Python", "Java", "Golang", "YAML", "C", "C++"}
	slice_2 := []int{23, 56, 21, 67, 32, 89, 45, 33}

	fmt.Println("The slices before sorting")
	fmt.Println("The first slice is: ", slice_1)
	fmt.Println("The second slice is: ", slice_2)

	sort.Strings(slice_1)
	sort.Ints(slice_2)

	fmt.Println("The slices after sorting!")
	fmt.Println("The first slice after sorting: ", slice_1)
	fmt.Println("The second slice after sorting: ", slice_2)
}
