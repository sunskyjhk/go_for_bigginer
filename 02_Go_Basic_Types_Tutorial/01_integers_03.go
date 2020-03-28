package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var myInt int8
	for i := 0; i < 129; i++ {
		myInt += 1
	}
	fmt.Println(myInt) // prints out -127
}
