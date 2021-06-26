// Program to show the Go Concurrency

package main

import (
	"fmt"
	"time"
)

func loopval(str string) {
	for w := 0; w < 10; w++ {
		time.Sleep(5 * time.Second)
		fmt.Println(str)
	}

}

func main() {

	// Calling GoRoutine

	go loopval("Hello Go Concurrency")

	loopval("I love Golang!")
}
