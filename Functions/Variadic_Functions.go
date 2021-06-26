// Program to show the Variadic function

package main

import "fmt"

// variadic function

func sum(numbers ...int) {
	fmt.Println(numbers, " ", " ")
	total := 0
	for _, numbers := range numbers {
		total += numbers
	}
	fmt.Println(total)
}

func main() {

	sum(1, 2)
	sum(4, 5, 6, 7)

	numbers := []int{1, 2, 3, 5, 6, 7, 9}
	sum(numbers...)

}
