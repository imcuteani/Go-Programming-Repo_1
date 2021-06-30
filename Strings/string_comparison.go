// Program to show string comparison in Golang

package main

import "fmt"

func main() {

	str1 := "Geeks"
	str2 := "Geek"
	str3 := "Nerd"
	str4 := "GeekforGeeks"
	str5 := "Ninja"

	//checking to see if the strings are equal or not

	result1 := str1 == str2
	result2 := str2 == str1
	result3 := str3 == str4
	result4 := str4 == str5
	result5 := str1 == str5

	fmt.Println("Result 1: ", result1)
	fmt.Println("Result 2: ", result2)
	fmt.Println("Result 3: ", result3)
	fmt.Println("Result 4: ", result4)
	fmt.Println("Result 5: ", result5)

	result6 := str1 != str2
	result7 := str2 != str3
	result8 := str3 != str4
	result9 := str4 != str5
	result10 := str1 != str5

	fmt.Println("Result 6: ", result6)
	fmt.Println("Result 7: ", result7)
	fmt.Println("Result 8: ", result8)
	fmt.Println("Result 9: ", result9)
	fmt.Println("Result 10: ", result10)

}
