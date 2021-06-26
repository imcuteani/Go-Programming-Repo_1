// Program to show anonymous Go Routine

package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 8; i++ {
		fmt.Println(i)
	}

	fmt.Println("This is the main function")

	go func() {
		fmt.Println("This is example of goroutine function")
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("This is the end of goroutine function")
}
