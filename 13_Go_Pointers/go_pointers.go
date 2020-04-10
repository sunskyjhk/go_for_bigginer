package main

import "fmt"

func myTestFunc(a *int) {
	*a += 3
	fmt.Println(*a)
}

func main() {

	// Introduction
	a := 2
	myTestFunc(&a)
	fmt.Println(a)

	// Defining Pointers
	var age *int

	fmt.Println(age)
	fmt.Println(&age)

	// Assigning Values to Pointers
	age = new(int)
	*age = 26

	fmt.Println(*age)
	fmt.Println(&age)
}
