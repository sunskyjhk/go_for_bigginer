package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	name := MyFunction("John", "Lee")
	fmt.Println(name)

	// we can assign the results to multiple variables
	// by defining their names in a comma separated list
	// like so:
	fullName, err := MyFunctionWithError("Henry", "Ford")
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

	// Challenge Result
	sum := Add(3, 5)
	fmt.Println(sum)
}

func MyFunction(firstName string, lastName string) string {
	fullname := firstName + " " + lastName
	return fullname
}

// multiple results from a function
func MyFunctionWithError(firstName string, lastName string) (string, error) {
	if len(firstName) <= 1 {
		return "", errors.New("First Name Must Be At Least 2 letters long")
	}

	// if the first name passes the check, we return the
	// concatenated names and a nil error value
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

// Challenge

func Add(num1 int, num2 int) int {
	return num1 + num2
}
