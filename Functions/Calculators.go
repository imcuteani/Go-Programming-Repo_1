package main

import (
	"fmt"
	"math"
)

// simple function to perform calculation

// Function addition
func add(a int, b int) int {
	sum := a + b
	fmt.Printf("The simple addition is %d\n", sum)
	return sum
}

// Function substraction

func subs(a int, b int) int {
	sub := a - b
	fmt.Printf("The substraction result is %d\n", sub)
	return sub
}

// Function multiplication

func mul(a int, b int) int {
	mul := a * b
	fmt.Printf("The multiplication result is %d\n", mul)
	return mul
}

// Function Division

func div(a int, b int) int {
	div := a / b
	fmt.Printf("The division result is %d\n", div)
	return div
}

// Function Pow

func pow(a float64, b float64) float64 {
	power := math.Pow(a, b)
	fmt.Printf("The power calculation is %f\n", power)
	return power
}

// Function Percentage

func Percentg(a float64, b float64) float64 {
	diff := a - b
	percent := (diff / a) * float64(100)
	fmt.Printf("The percentage calculation of the numbers are %f\n", percent)
	return percent
}

// Function exponential

func expon(a float64) float64 {
	exponen := math.Exp(a)
	fmt.Printf("The exponential value is %f.3\n", exponen)
	return exponen

}

// Function Absolute value

func abs(a float64) float64 {
	abs := math.Abs(a)
	fmt.Printf("The absolute value is %f\n", abs)
	return abs
}

// Function Ceil value

func ceil(a float64) float64 {
	ceilval := math.Ceil(a)
	fmt.Printf("The ceiling value is %f\n", ceilval)
	return ceilval
}

// Function floor value

func floor(a float64) float64 {
	floorval := math.Floor(a)
	fmt.Printf("The floor value is %f\n", floorval)
	return floorval

}

// Logarithm in base of 10 function

func logbase10(a float64) float64 {
	log10val := math.Log10(a)
	fmt.Printf("The log 10 base value of the number is %f\n", log10val)
	return log10val
}

// Function to return maximum between two numbers

func max(a float64, b float64) float64 {
	Maxnum := math.Max(a, b)
	fmt.Printf("The max between two numbers are %f\n", Maxnum)
	return Maxnum
}

// Function to return between two numbers

func min(a float64, b float64) float64 {
	Minnum := math.Min(a, b)
	fmt.Printf("The minimum between two numbers are %f\n", Minnum)
	return Minnum
}

// Function to return Mod value

func mod(a float64, b float64) float64 {
	Modulus := math.Mod(a, b)
	fmt.Printf("The modulus of these two numbers are: %f\n", Modulus)
	return Modulus
}

// Main method

func main() {
	fmt.Println("The calculator program in Go!")
	add(10, 8)
	subs(200, 109)
	mul(301, 261)
	div(105, 15)
	pow(6.1, 3.5)
	Percentg(2.5, 201.5)
	expon(6.66)
	abs(9.1)
	ceil(101.3)
	floor(101.3)
	mod(51.3, 43.3)
	max(109.45, 99.4)
	min(351.20, 771.1)

	fmt.Println("The end of the calculator program in Go!")
}
