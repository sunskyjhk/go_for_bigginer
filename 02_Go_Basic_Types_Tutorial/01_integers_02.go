package main

import "fmt"

var myInt int8

func main() {
	myInt = 2500
	fmt.Println(myInt)
}

// excute this file
// go run ./Go_Basic_Types_Tutorial/integers_02.go
//
// Print
// # command-line-arguments
// Go_Basic_Types_Tutorial/integers_02.go:8:8: constant 2500 overflows int8
