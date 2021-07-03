// Program to show Method definitions in Golang

package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perimeter() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 4}

	fmt.Println("The area is:", r.area())
	fmt.Println("The perimeter is: ", r.perimeter())

	rp := &r
	fmt.Println("The area is:", rp.area())
	fmt.Println("The Perimeter is:", rp.perimeter())
}
