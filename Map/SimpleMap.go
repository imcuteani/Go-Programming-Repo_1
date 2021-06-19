// Program to show the Simple Map Programming with Go

package main

import "fmt"

func main() {

	//Create a simple Map by using Make function along with Key & value pairs

	mapval := make(map[string]int)

	mapval["k1"] = 7
	mapval["k2"] = 15

	val1 := mapval["k1"]

	fmt.Println("Map is: ", val1)

	fmt.Println("the length is: ", len(mapval))

	delete(mapval, "k2")
	fmt.Println("Map: ", mapval)

	_, prs := mapval["k2"]
	fmt.Println("prs: ", prs)

	// Another way to initialize Map

	mapval2 := map[string]int{"val1": 1, "val2": 2}
	fmt.Println("The second Map value is: ", mapval2)

}
