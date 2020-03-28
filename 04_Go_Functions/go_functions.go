package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	name := myfunction("John", "Lee")
	fmt.Println(name)

	// we can assign the results to multiple variables
	// by defining their names in a comma separated list
	// like so:
	fullName, err := myfunction2("Henry", "Ford")
	if err != nil {
		fmt.Println("Handle Error Case")
	}
	fmt.Println(fullName)

	// anonymous function
	myFunc := addOne()
	fmt.Println(myFunc()) // 2
	fmt.Println(myFunc()) // 3
	fmt.Println(myFunc()) // 4
	fmt.Println(myFunc()) // 5
}

func myfunction(firstName string, lastName string) string {
	fullname := firstName + " " + lastName
	return fullname
}

// multiple results from a function
func myfunction2(firstName string, lastName string) (string, error) {
	return firstName + " " + lastName, nil
}

// anonymous functions
func addOne() func() int {
	var x int
	// we define and return an
	// anonymous function which in turn
	// returns an integer value
	return func() int {
		// this anonymous function
		// has access to the x variable
		// defined in the parent function
		x++
		return x + 1
	}
}
