// how to run this file:
// go run example_01.go

package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
