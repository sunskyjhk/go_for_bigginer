package main

import "fmt"

func main() {
	var isTrue bool = true
	var isFalse bool = false

	// AND
	if isTrue && isFalse {
		fmt.Println("Both conditions need to be True")
	}
	// OR - not exclusive
	if isTrue || isFalse {
		fmt.Println("Only one condition needs to be True")
	}
}
