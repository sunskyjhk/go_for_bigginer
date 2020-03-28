package main

import "fmt"

func main() {
	var men uint8
	men = 5
	var women int16
	women = 6

	var people int
	// this throws a compile error
	// -> people = men + women
	// this handles converting to a standard format
	// and is legal within your go programs
	people = int(men) + int(women)
	fmt.Println(people)
}
