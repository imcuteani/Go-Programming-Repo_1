// Go program to show the pointer to Struct

package main

import "fmt"

type Car struct {
	manufacturer, Type, color string
	model, number, maxspeed   int
}

func main() {

	c := &Car{"BMW", "SUV", "Maize", 301, 001, 140}

	fmt.Println("The Manufacturer is: ", (*c).manufacturer)
	fmt.Println("The Car type is: ", (*c).Type)
	fmt.Println("The Car Color is: ", (*c).color)
	fmt.Println("The MaxSpeed of the car is: ", (*c).maxspeed)

}
